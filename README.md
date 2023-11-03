# strava-snowflake
track and export personal Strava data.  Ingest and visualize in Snowflake.

copy github.com/stevenpelley/strava3golang/oauth_client_config.json and the resulting token.json here.
Also needs snowflake key pair credentials placed in snowflake_config (see template)

state:
passing tests for both duckdb and snowflake
downloads activities within time range from strava
skips retrieving streams for activities already in database or listed in command line options
stores and transforms data in duckdb.
python snowpark udfs organized, flatten_streams written and tested.  Additionally fills time gaps.

todo:
snowpark for data cleaning and bike stress
- create CLI command to register udfs
- create exponentially weighted average udf

use terraform for reproducible snowflake and schema creation.
try out snowflake builtin dashboarding
try out some other dashboard options (superset, mesosphere)
deploy as a stored procedure using new external network integration.
set up a frontend for web hooks and use tasks to fetch new data immediately
investigate dbt?