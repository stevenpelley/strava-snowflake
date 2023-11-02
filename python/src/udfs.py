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
        types.StructField('inserted_for_gap', types.BooleanType()),
    ])
    stream_names = list([s.lower() for s in struct_type.names])[0:-1]

    time_idx = stream_names.index("time")
    watts_idx = stream_names.index("watts")
    inserted_for_gap_idx = len(stream_names)

    @staticmethod
    def _registration_kwargs(is_permanent=False):
        kwargs = {
            "output_schema": FlattenStreams.struct_type,
            "input_types": [types.VariantType(), types.BooleanType(), types.IntegerType()],
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

    @staticmethod
    def fill_gaps(
            it: Iterable[tuple],
            seconds_to_add_at_end: int):
        """
        fill any gaps in time and add time to the end for calculating EWA.
        gaps in time are filled with 0 watts of power and "inserted_for_gap" = True.
        All other columns receive None.

        Assumes that rows appear with times in monotonic increasing order.

        Returns an iterator of the full stream with newly inserted rows.
        """

        # set up by asserting initial state and storing the first tuple's time.
        try:
            t = next(it)
        except StopIteration:
            raise Exception("FlattenStream.fill_gaps input iterator requires at least 1 item")
        assert t[FlattenStreams.time_idx] == 0, (
            "FlattenStream.fill_gaps first tuple's time must be 0.  Found: {}"
            ).format(t[FlattenStreams.time_idx])
        previous_time = t[FlattenStreams.time_idx]
        yield t

        # for each new row pad any gap from the last row before yielding the row
        try:
            while True:
                t = next(it)
                this_time = t[FlattenStreams.time_idx]
                assert this_time > previous_time, (
                    "expecting time to increase row-to-row.  Previous: {}, This: {}"
                    ).format(previous_time, this_time)
                yield from FlattenStreams.pad_zero_watts(previous_time, this_time)

                previous_time = this_time
                yield t
        except StopIteration:
            # now pad the end
            yield from FlattenStreams.pad_zero_watts(
                previous_time,
                previous_time + seconds_to_add_at_end + 1)
    
    @staticmethod
    def pad_zero_watts(
            previous_time: int,
            next_time: int,):
        """
        Yield tuples with times and 0 watts to pad a gap between previous_time
        and next_time.  This does not introduce any rows for either previous_time
        or next_time (the range is exclusive at both ends).
        The yielded tuples have inserted_for_gap=True.  All columns other than
        time, watts, and inserted_for_gap will be None.
        """
        for time in range(previous_time + 1, next_time):
            l = [None] * len(FlattenStreams.struct_type.fields)
            l[FlattenStreams.time_idx] = time
            l[FlattenStreams.watts_idx] = 0
            l[FlattenStreams.inserted_for_gap_idx] = True
            yield tuple(l)
        
    def process(
            self,
            data: dict,
            fill_gaps: bool,
            seconds_to_add_at_end: int) -> Iterable[tuple[
            int, int, int, int, float, float, float, bool, float, float, int, float, bool]]:
        ss = data['StreamSet']
        assert 'time' in ss, "stream 'time' must exist"
        default_length = len(ss['time']['data'])

        iters = []
        for stream in self.stream_names:
            iters.append(self.iter_for_stream(
                ss,
                self.use_stream_name.get(stream, stream),
                default_length,
                self.mappers.get(stream, lambda x: x)))
        # add "inserted_for_gap" column
        iters.append(itertools.repeat(False, default_length))

        zipped = zip(*iters, strict=True)
        if fill_gaps:
            zipped = FlattenStreams.fill_gaps(zipped, seconds_to_add_at_end)
        yield from zipped