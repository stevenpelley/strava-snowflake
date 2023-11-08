import src.connect
import src.ewma_udf

from snowflake.snowpark import Session
from snowflake.snowpark import udtf 
from snowflake.snowpark import functions

import numpy
import pandas
import typing


def test_sf_ewma() -> None:
    with src.connect.create_session() as session:
        sf_create_test_data(session)
        # anonymous
        udf = src.ewma_udf.EWMA.register(session)
        sf_test_ewma(session, udf)

def get_test_data() -> list[tuple[int, int, int]]:
    return [
        (0,0,50),
        (0,1,100),
        (0,2,200),
        (0,3,150),
        (0,4,50),
        (1,0,100),
        (1,3,100),
        (1,4,100),
        (1,5,100),
        (1,10,100),
    ]

def assert_test_result(result):
    r = [t[0:2] + (int(round(t[2])),) for t in result]
    assert r == [
        (0,0,50),
        (0,1,75),
        (0,2,118),
        (0,3,126),
        (0,4,110),
        (1,0,100),
        (1,3,100),
        (1,4,100),
        (1,5,100),
        (1,10,100),
    ]

def sf_create_test_data(session:Session) -> None:
    session.connection.cursor().execute(
        "create temp table strava.activities_test.etl_test (etl_id number, time number, watts number);")

    td = get_test_data()
    session.connection.cursor().execute(
        "insert into strava.activities_test.etl_test values (?,?,?)",
        [
            [t[0] for t in td],
            [t[1] for t in td],
            [t[2] for t in td],
        ])

def sf_test_ewma(session: Session, name_or_udf: typing.Union[str, udtf.UserDefinedTableFunction]):
    if isinstance(name_or_udf, str):
        my_udtf = functions.table_function(name_or_udf)
    elif isinstance(name_or_udf, udtf.UserDefinedTableFunction):
        my_udtf = name_or_udf
    else:
        raise Exception("unexpected type: {}".format(name_or_udf.__class__.__name__))

    df1 = session.table(['STRAVA', 'ACTIVITIES_TEST', 'ETL_TEST'])
    df2 = df1.join_table_function(
            my_udtf(
                df1['TIME'],
                df1['WATTS']
            ).over(
                partition_by=df1['ETL_ID'],
                order_by=df1['TIME']
            ).as_(
                "EWMA_TIME",
                "EWMA",
            )
        ).select(
            functions.col('ETL_ID'),
            functions.col('EWMA_TIME'),
            functions.col('EWMA'),
        ).sort(
            functions.col('ETL_ID'),
            functions.col('EWMA_TIME')
        ).collect()
    assert_test_result(df2)

def test_ewa_pandas() -> None:
    #df = pandas.DataFrame(
    #    {"val": numpy.concatenate([
    #        numpy.repeat(0.0, 20),
    #        numpy.repeat(1.0, 10)])})
    #df = df.ewm(
    #    halflife=2.0,
    #    #alpha=1.0-math.exp(3.0/4.0),
    #    #adjust=False,
    #    ).mean()
    #print(df.to_string())

    # conclusion: just use halflife over the number of samples or time period.  It works.
    # don't bother with adjust, take Coggan's advice and throw out the first and last 30 seconds
    # (post-ewma).  Don't worry about padding the end either.

    # now let's test the axis.  Do we need to fill in missing values?

    df = pandas.DataFrame(
        {"val": numpy.concatenate([
            numpy.repeat(0.0, 5),
            numpy.repeat(1.0, 5)])},
        index=[0, 1, 6, 8, 9, 10, 14, 16, 18, 19])
    df = df.reindex(pandas.RangeIndex(20))
    df = df.ewm(
        halflife=2.0,
        #alpha=1.0-math.exp(3.0/4.0),
        adjust=False,
        ignore_na=False,
        ).mean()
    print(df.to_string())

    # conclusion: if we want to treat pause gaps as 0s we can reindex the dataframe within
    # the UDTF and then "don't ignore N/As".  This avoids having to pad the input