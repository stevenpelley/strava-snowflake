import os
import json

from snowflake.snowpark import functions
from snowflake.snowpark import Session
from snowflake.snowpark import types

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

def run(session):
    sess.udf()
    df1 = sess.sql('select count(*) from strava.activities.etl')
    df1.show()

def udf_id(session):
    def udf_def(data: dict) -> int:
        return data['Activity']['id']
    my_udf = sess.udf.register(udf_def)
    sess.table(['STRAVA', 'ACTIVITIES', 'ETL']).select(my_udf(functions.col('DATA'))).limit(5).show()

def udtf_name_and_id(session):
    @functions.udtf(output_schema=['NAME', 'ID'], input_types=[types.VariantType()])
    class udtf(object):
        def process(self, data) -> Iterable[tuple[str, int]]:
            activity = data['Activity']
            yield (activity['name'], activity['id'],)
    sess.table(['STRAVA', 'ACTIVITIES', 'ETL']).select(udtf(functions.col('DATA'))).limit(5).show()

def udtf_flatten_time(session):
    @functions.udtf(output_schema=['TIME'], input_types=[types.VariantType()])
    class udtf(object):
        def process(self, data) -> Iterable[tuple[int]]:
            l = data['StreamSet']['time']['data']
            for item in l:
                yield (item,)
    the_count = sess.table(['STRAVA', 'ACTIVITIES', 'ETL']).select(udtf(functions.col('DATA'))).count()
    print(the_count)

def udtf_flatten_all(session):
    class udtf(object):
        @staticmethod
        def iter_for_stream(streamset, name, default_length):
            stream = streamset.get(name)
            if stream:
                return (stream['data'],)
            else:
                return (itertools.repeat(None, default_length),)

        @staticmethod
        def iter_for_latlng(streamset, default_length):
            stream = streamset.get('latlng')
            if stream:
                (lat, lng) = latlng_iter = itertools.tee(stream['data'], 2)
                lat = map(operator.itemgetter(0), lat)
                lng = map(operator.itemgetter(1), lng)
                return (lat, lng,)
            else:
                return (itertools.repeat(None, default_length),
                        itertools.repeat(None, default_length),)
            
            
        def process(self, data) -> Iterable[tuple[int]]:
            ss = data['StreamSet']

            default_length = len(ss['time']['data'])
            zipped = zip(*itertools.chain(
                self.iter_for_stream(ss, 'time', default_length),
                self.iter_for_stream(ss, 'watts', default_length),
                self.iter_for_stream(ss, 'heartrate', default_length),
                self.iter_for_stream(ss, 'cadence', default_length),
                self.iter_for_stream(ss, 'velocity_smooth', default_length),
                self.iter_for_stream(ss, 'grade_smooth', default_length),
                self.iter_for_stream(ss, 'distance', default_length),
                self.iter_for_stream(ss, 'moving', default_length),
                self.iter_for_latlng(ss, default_length),
                self.iter_for_stream(ss, 'temp', default_length),
                self.iter_for_stream(ss, 'altitude', default_length),
                ), strict=True)
            for tup in zipped:
                yield tup

    my_udtf = sess.udtf.register(udtf, output_schema=streams_struct_type, input_types=[types.VariantType()])
    the_count = sess.table(['STRAVA', 'ACTIVITIES', 'ETL']
                          ).filter(functions.col('ETL_ID') == 444
                          ).select(functions.col('ETL_ID'), my_udtf(functions.col('DATA'))
                          ).limit(10
                          ).show()
    
if __name__ == "__main__":
    sess = create_session()
    #run(sess)
    #udf_id(sess)
    #udtf_name_and_id(sess)
    #udtf_flatten_time(sess)
    udtf_flatten_all(sess)