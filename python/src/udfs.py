from collections.abc import Iterable
import itertools
import operator

from snowflake.snowpark import types
from snowflake.snowpark import Session


class FlattenStreams(object):
    struct_type = types.StructType([
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
    output_column_names = list([s.lower() for s in struct_type.names])

    @staticmethod
    def _registration_kwargs(is_permanent=False):
        kwargs = {
            "output_schema": FlattenStreams.struct_type,
            "input_types": [types.VariantType()],
        }
        if is_permanent:
            kwargs.update({
                "name" : ["STRAVA", "UDFS", "FLATTEN_STREAMS"],
                "is_permanent" : True,
                "stage_location" : '"STRAVA"."UDFS"."UDF_STAGE"',
                "replace" : True,
            })
        return kwargs

    @staticmethod
    def register(session: Session, is_permanent=False):
        return session.udtf.register(
            FlattenStreams,
            **FlattenStreams._registration_kwargs(is_permanent))


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
        for stream in self.output_column_names:
            iters.append(self.iter_for_stream(
                ss,
                self.use_stream_name.get(stream, stream),
                default_length,
                self.mappers.get(stream, lambda x: x)))
        zipped = zip(*iters, strict=True)
        for tup in zipped:
            yield tup