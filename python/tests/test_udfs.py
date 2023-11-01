import src.connect
import src.udfs

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
                          ).select(functions.col('ETL_ID'), my_udtf(functions.col('DATA'))
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

def test_unit_flatten_streams():
    fs = src.udfs.FlattenStreams()

    obj = {"StreamSet": {"time" : {"data" : [1, 2, 3]}, "watts" : {"data" : [100, 101, 102]}}}
    result = list(fs.process(obj))
    expected = [(1, 100,) + ((None,) * 10), (2, 101,) + ((None,) * 10), (3, 102,) + ((None,) * 10)]
    assert result == expected, "expected: {}, found: {}".format(expected, result)

    obj['StreamSet']['latlng'] = {"data": [[10, 20], [11, 21], [12, 22]]}
    result = list(fs.process(obj))
    expected2 = [list(e) for e in expected]
    expected2[0][8] = 10
    expected2[0][9] = 20
    expected2[1][8] = 11
    expected2[1][9] = 21
    expected2[2][8] = 12
    expected2[2][9] = 22
    expected2 = [tuple(l) for l in expected2]
    assert result == expected2, "expected: {}, found: {}".format(expected2, result)