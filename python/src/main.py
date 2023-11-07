from src.connect import create_session
from src.flatten_streams_udf import FlattenStreams
from src.ewma_udf_nonvectorized import EWMA as EWMA_NV
from src.ewma_udf import EWMA
import snowflake.snowpark

def register_udtfs(session: snowflake.snowpark.Session):
    FlattenStreams.register(session, is_permanent=True)
    EWMA.register(session, is_permanent=True)
    EWMA_NV.register(session, is_permanent=True)

if __name__=="__main__":
    session = create_session()
    register_udtfs(session)