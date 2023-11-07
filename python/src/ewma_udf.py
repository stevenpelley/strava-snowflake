import pandas
import typing

from snowflake.snowpark import types
from snowflake.snowpark import Session


class EWMA(object):
    output_schema = types.PandasDataFrameType(
        col_types=[types.IntegerType(), types.DoubleType()],
        col_names=['TIME', 'EWMA'])

    registered_name_tup = ["STRAVA", "UDFS", "EWMA"]

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

    # note: cannot (yet?) indicate type hint for return type of
    # python vectorized udtf
    # note: we return the time so that the calculated ewma value can be correlated
    # against it.  UDTF output is correlated against the input only when using
    # the process() method, which by definition is not available for vectorized UDTFs
    def end_partition(self, df: types.PandasDataFrame):
        # pull out all the stops, force my IDE to know this is a DataFrame
        df: pandas.DataFrame = typing.cast(pandas.DataFrame, df) # type: pandas.DataFrame
        df2 = df['WATTS']
        ewma = df2.ewm(halflife=30.0).mean()
        return pandas.DataFrame(data = {
            "EWMA_TIME": df['TIME'],
            "EWMA": ewma
        })

EWMA.end_partition._sf_vectorized_input = pandas.DataFrame