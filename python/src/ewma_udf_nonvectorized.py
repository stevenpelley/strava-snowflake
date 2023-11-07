import math
from collections.abc import Iterable

from snowflake.snowpark import types
from snowflake.snowpark import Session


class EWMA(object):
    output_schema = types.StructType([
        types.StructField('EWMA', types.DoubleType()),
    ])

    registered_name_tup = ["STRAVA", "UDFS", "EWMA_NONVECTORIZED"]

    @staticmethod
    def _registration_kwargs(is_permanent=False):
        kwargs = {
            "output_schema": EWMA.output_schema,
            "input_types": [types.IntegerType(), types.IntegerType()],
            "input_names": ["TIME", "WATTS"],
        }
        if is_permanent:
            kwargs.update({
                "name" : EWMA.registered_name_tup,
                "is_permanent" : True,
                "stage_location" : '"STRAVA"."UDFS"."UDF_STAGE"',
                "replace" : True,
            })
        return kwargs

    @staticmethod
    def register(session: Session, is_permanent=False):
        return session.udtf.register(
            EWMA,
            **EWMA._registration_kwargs(is_permanent))

    def __init__(self):
        self.accum = 0.0
        halflife = 30.0
        self.alpha = 1 - math.exp(-math.log(2.0)/halflife)
        self.one_minus_alpha = 1.0 - self.alpha

    def process(
            self,
            time: int, # here in case we want to handle gaps
            watts: int) -> Iterable[tuple[float]]:
        self.accum = (self.one_minus_alpha * self.accum) + (self.alpha * watts)
        yield (self.accum,)
