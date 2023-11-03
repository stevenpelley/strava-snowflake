import src.connect
import src.udfs

import copy
import json
import pytest
import operator
import typing

from snowflake.snowpark import functions
from snowflake.snowpark import Session
from snowflake.snowpark import udtf

def sf_create_test_data(session:Session) -> None:
    session.connection.cursor().execute(
        "create temp table strava.activities_test.etl_test (etl_id number, data variant);")

    base_obj = {"Activity": {"id": 0}, "StreamSet": {"time": {"data": []}, "watts": {"data": []}}}
    obj1 = copy.deepcopy(base_obj)
    obj1['Activity']['id'] = 100
    obj1['StreamSet']['time']['data'].extend([0, 1, 2, 3])
    obj1['StreamSet']['watts']['data'].extend([10, 20, 30, 40])
    obj2 = copy.deepcopy(base_obj)
    obj1['Activity']['id'] = 200
    obj2['StreamSet']['time']['data'].extend([0, 1, 2, 10, 11])
    obj2['StreamSet']['watts']['data'].extend([50, 50, 0, 0, 40])

    session.connection.cursor().execute(
        ("insert into strava.activities_test.etl_test select $1, "
                "parse_json($2) from values (1, '{}'), (2, '{}')").format(
        json.dumps(obj1), json.dumps(obj2)))

def sf_test_flatten_streams(session: Session, name_or_udf: typing.Union[str, udtf.UserDefinedTableFunction]):
    kwargs = {}
    if isinstance(name_or_udf, str):
        my_udtf = functions.table_function(name_or_udf)
    elif isinstance(name_or_udf, udtf.UserDefinedTableFunction):
        my_udtf = name_or_udf
    else:
        raise Exception("unexpected type: {}".format(name_or_udf.__class__.__name__))

    results = session.table(['STRAVA', 'ACTIVITIES_TEST', 'ETL_TEST']
                          ).select(
                              functions.col('ETL_ID'),
                              my_udtf(
                                  functions.col('DATA'),
                                  functions.lit(True),
                                  functions.lit(2))
                          ).sort(functions.col('etl_id'), functions.col('time')).collect()

    # look only at etl_id, time, watts, heartrate, and inserted_for_gap
    results = list(map(
        operator.itemgetter('ETL_ID','TIME','WATTS','HEARTRATE','INSERTED_FOR_GAP'),
        results))
    assert results == [
        (1, 0, 10, None, False,),
        (1, 1, 20, None, False,),
        (1, 2, 30, None, False,),
        (1, 3, 40, None, False,),
        (1, 4, 0, None, True,),
        (1, 5, 0, None, True,),
        (2, 0, 50, None, False,),
        (2, 1, 50, None, False,),
        (2, 2, 0, None, False,),
        (2, 3, 0, None, True,),
        (2, 4, 0, None, True,),
        (2, 5, 0, None, True,),
        (2, 6, 0, None, True,),
        (2, 7, 0, None, True,),
        (2, 8, 0, None, True,),
        (2, 9, 0, None, True,),
        (2, 10, 0, None, False,),
        (2, 11, 40, None, False,),
        (2, 12, 0, None, True,),
        (2, 13, 0, None, True,),
    ]

def test_sf_flatten_streams():
    session = src.connect.create_session()
    sf_create_test_data(session)

    # anonymous
    udf = src.udfs.FlattenStreams.register(session)
    sf_test_flatten_streams(session, udf)
    
    # named
    # TODO: delete the function and associated stage files.  Right now we overwrite them
    # each time so size should never grow
    udf = src.udfs.FlattenStreams.register(session, is_permanent=True)
    # passing a string calls the named function
    sf_test_flatten_streams(session, ".".join(src.udfs.FlattenStreams.registered_name_tup))

def update_list(l, tups):
    for (idx, val,) in tups:
        l[idx] = val
    return l

def test_unit_flatten_streams():
    fs = src.udfs.FlattenStreams()

    obj = {"StreamSet": {"time" : {"data" : [0, 1, 2]}, "watts" : {"data" : [100, 101, 102]}}}
    result = list(fs.process(obj, False, 0))
    empty_row = (None,) * 13
    expected = [
        tuple(update_list(list(empty_row), [(0, 0,), (1, 100,), (12, False)])),
        tuple(update_list(list(empty_row), [(0, 1,), (1, 101,), (12, False)])),
        tuple(update_list(list(empty_row), [(0, 2,), (1, 102,), (12, False)])),
        ]
    assert result == expected, "expected: {}, found: {}".format(expected, result)

    obj['StreamSet']['latlng'] = {"data": [[10, 20], [11, 21], [12, 22]]}
    result = list(fs.process(obj, False, 0))
    expected2 = [list(e) for e in expected]
    expected2[0][8] = 10
    expected2[0][9] = 20
    expected2[1][8] = 11
    expected2[1][9] = 21
    expected2[2][8] = 12
    expected2[2][9] = 22
    expected2 = [tuple(l) for l in expected2]
    assert result == expected2, "expected: {}, found: {}".format(expected2, result)

def test_unit_flatten_streams_with_gaps():
    fs = src.udfs.FlattenStreams()

    obj = {"StreamSet": {
        "time" : {"data" : [0, 1, 5]},
        "watts" : {"data" : [100, 101, 102]},
        "heartrate" : {"data" : [10, 11, 12]},
        }}
    result = list(fs.process(obj, True, 2))
    empty_row = (None,) * 13
    expected = [
        tuple(update_list(list(empty_row), [(0, 0,), (1, 100,), (2, 10,), (12, False)])),
        tuple(update_list(list(empty_row), [(0, 1,), (1, 101,), (2, 11,), (12, False)])),
        tuple(update_list(list(empty_row), [(0, 2,), (1, 0,), (12, True)])),
        tuple(update_list(list(empty_row), [(0, 3,), (1, 0,), (12, True)])),
        tuple(update_list(list(empty_row), [(0, 4,), (1, 0,), (12, True)])),
        tuple(update_list(list(empty_row), [(0, 5,), (1, 102,), (2, 12,), (12, False)])),
        tuple(update_list(list(empty_row), [(0, 6,), (1, 0,), (12, True)])),
        tuple(update_list(list(empty_row), [(0, 7,), (1, 0,), (12, True)])),
        ]
    assert result == expected, "expected: {}, found: {}".format(expected, result)

def test_unit_flatten_streams_errors():
    # must have non-empty stream data
    with pytest.raises(
            Exception,
            match="FlattenStream.fill_gaps input iterator requires at least 1 item"):
        fs = src.udfs.FlattenStreams()
        obj = {"StreamSet": {"time" : {"data" : []}}}
        result = list(fs.process(obj, True, 0))

    # time stream must exist
    with pytest.raises(
            AssertionError,
            match="stream 'time' must exist"):
        fs = src.udfs.FlattenStreams()
        obj = {"StreamSet": {"watts" : {"data" : [100]}}}
        result = list(fs.process(obj, True, 0))

    # streams must be same length
    with pytest.raises(
            ValueError,
            match="zip\(\) argument [0-9]+ is shorter than argument [0-9]+"):
        fs = src.udfs.FlattenStreams()
        obj = {"StreamSet": {"time" : {"data" : [0, 2, 1]}, "watts": {"data" : [100]}}}
        result = list(fs.process(obj, True, 0))
    
    # first time must be 0
    with pytest.raises(
            AssertionError,
            match="FlattenStream.fill_gaps first tuple's time must be 0.  Found: 2"):
        fs = src.udfs.FlattenStreams()
        obj = {"StreamSet": {"time" : {"data" : [2]}}}
        result = list(fs.process(obj, True, 0))
    
    # times must increase monotonically
    #                "expecting time to increase row-to-row.  Previous: {}, This: {}"
    with pytest.raises(
            AssertionError,
            match="expecting time to increase row-to-row.  Previous: 2, This: 1"):
        fs = src.udfs.FlattenStreams()
        obj = {"StreamSet": {"time" : {"data" : [0, 2, 1]}}}
        result = list(fs.process(obj, True, 0))