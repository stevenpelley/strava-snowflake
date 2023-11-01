import os
import json
import typing

from snowflake.snowpark import functions
from snowflake.snowpark import Session
from snowflake.snowpark import types
from snowflake.snowpark import udtf

from collections.abc import Iterable
from collections import namedtuple
import itertools
import operator

from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives.asymmetric import dsa
from cryptography.hazmat.primitives import serialization

streams_struct_type = types.StructType([
    types.StructField('time', types.LongType()),
    types.StructField('watts', types.LongType()),
    types.StructField('heartrate', types.LongType()),
    types.StructField('cadence', types.LongType()),
    types.StructField('velocity_smooth', types.DoubleType()),
    types.StructField('grade_smooth', types.DoubleType()),
    types.StructField('distance', types.DoubleType()),
    types.StructField('moving', types.BooleanType()),
    types.StructField('lat', types.DoubleType()),
    types.StructField('lng', types.DoubleType()),
    types.StructField('temp', types.LongType()),
    types.StructField('altitude', types.DoubleType()),
])
output_column_names = list([s.lower() for s in streams_struct_type.names])

def create_session():
    # run from project root
    with open("snowflake_config.json", "r") as f:
        j = json.loads(f.read())

    with open(j['private_key_file'], "rb") as key:
        p_key= serialization.load_pem_private_key(
            key.read(),
            password=None,
            backend=default_backend()
        )

    pkb = p_key.private_bytes(
        encoding=serialization.Encoding.DER,
        format=serialization.PrivateFormat.PKCS8,
        encryption_algorithm=serialization.NoEncryption())

    return Session.builder.configs(
        {
            "user":j['user'],
            "account":j['account'],
            "private_key":pkb,
        }).create()

class FlattenStreams(object):
    def __init__(self):
        self.mappers = {"lat": operator.itemgetter(0), "lng": operator.itemgetter(1)}
        self.use_stream_name = {"lat": "latlng", "lng": "latlng"}

    @staticmethod
    def iter_for_stream(streamset, name, default_length, mapper):
        stream = streamset.get(name)
        if stream:
            return map(mapper, stream['data'])
        else:
            return itertools.repeat(None, default_length)
        
    def process(self, data) -> Iterable[tuple[int]]:
        ss = data['StreamSet']
        assert 'time' in ss
        default_length = len(ss['time']['data'])

        iters = []
        for stream in output_column_names:
            iters.append(self.iter_for_stream(
                ss,
                self.use_stream_name.get(stream, stream),
                default_length,
                self.mappers.get(stream, lambda x: x)))
        zipped = zip(*iters, strict=True)
        for tup in zipped:
            yield tup

def register_flatten_streams(session, is_permanent=False):
    kwargs = {}
    if is_permanent:
        kwargs = {
            "name" : ["STRAVA", "UDFS", "FLATTEN_STREAMS"],
            "is_permanent" : True,
            "stage_location" : '"STRAVA"."UDFS"."UDF_STAGE"',
            "replace" : True,
        }
        
    return sess.udtf.register(
        FlattenStreams,
        output_schema=streams_struct_type,
        input_types=[types.VariantType()],
        **kwargs)

def test_flatten_streams(name_or_udf: typing.Union[str, udtf.UserDefinedTableFunction]):
    kwargs = {}
    if isinstance(name_or_udf, str):
        my_udtf = functions.table_function(name_or_udf)
    elif isinstance(name_or_udf, udtf.UserDefinedTableFunction):
        my_udtf = name_or_udf
    else:
        raise Exception("unexpected type: {}".format(name_or_udf.__class__.__name__))

    sess.table(['STRAVA', 'ACTIVITIES', 'ETL']
                          ).filter(functions.col('ETL_ID') == 444
                          ).select(functions.col('ETL_ID'), my_udtf(functions.col('DATA'))
                          ).limit(10
                          ).show()

def local_test_flatten_streams():
    fs = FlattenStreams()

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
    
if __name__ == "__main__":
    sess = create_session()

    # something like a unit test
    #local_test_flatten_streams()

    # anonymous udtf
    #my_udtf = register_flatten_streams(sess)
    #test_flatten_streams(my_udtf)

    # named udtf
    register_flatten_streams(sess, True)
    test_flatten_streams('"STRAVA"."UDFS"."FLATTEN_STREAMS"')
