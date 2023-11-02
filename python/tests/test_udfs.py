import src.connect
import src.udfs

import pytest
import typing

from snowflake.snowpark import functions
from snowflake.snowpark import Session
from snowflake.snowpark import udtf

def sf_test_flatten_streams(session: Session, name_or_udf: typing.Union[str, udtf.UserDefinedTableFunction]):
    kwargs = {}
    if isinstance(name_or_udf, str):
        my_udtf = functions.table_function(name_or_udf)
    elif isinstance(name_or_udf, udtf.UserDefinedTableFunction):
        my_udtf = name_or_udf
    else:
        raise Exception("unexpected type: {}".format(name_or_udf.__class__.__name__))

    c = session.table(['STRAVA', 'ACTIVITIES', 'ETL']
                          ).filter(functions.col('ETL_ID') == 444
                          ).select(
                              functions.col('ETL_ID'),
                              my_udtf(
                                  functions.col('DATA'),
                                  functions.lit(True),
                                  functions.lit(2))
                          ).limit(10
                          ).count()
    assert c == 10, "expected 10, found: {}".format(c)

def test_sf_flatten_streams():
    # anonymous
    session = src.connect.create_session()
    udf = src.udfs.FlattenStreams.register(session)
    sf_test_flatten_streams(session, udf)
    
    #named
    pass

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