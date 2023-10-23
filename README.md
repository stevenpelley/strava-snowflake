# strava-snowflake
track and export personal Strava data.  Ingest and visualize in Snowflake.

copy github.com/stevenpelley/strava3golang/oauth_client_config.json and the resulting token.json here.

state:
downloads activities within time range from strava
skips retrieving streams for activities already in database or listed in command line options
stores and transforms data in duckdb.
connection demo working in snowflake

todo:
update to load data into snowflake
snowpark for data cleaning and bike stress
try out snowflake builtin dashboarding
try out some other dashboard options (superset, mesosphere)
investigate dbt