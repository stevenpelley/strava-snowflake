# strava-snowflake
track and export personal Strava data.  Ingest and visualize in Snowflake.

copy github.com/stevenpelley/strava3golang/oauth_client_config.json and the resulting token.json here.
Also needs snowflake key pair credentials placed in snowflake_config (see template)

state:
snowflake passes simple tests
downloads activities within time range from strava
skips retrieving streams for activities already in database or listed in command line options
stores and transforms data in duckdb.
connection demo working in snowflake

todo:
update activitiessql cmd to use snowflake
snowpark for data cleaning and bike stress
use terraform for reproducible snowflake and schema creation.
try out snowflake builtin dashboarding
try out some other dashboard options (superset, mesosphere)
deploy as a stored procedure using new external network integration.
set up a frontend for web hooks and use tasks to fetch new data immediately
investigate dbt?