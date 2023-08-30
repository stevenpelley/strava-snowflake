# strava-snowflake
track and export personal Strava data.  Ingest and visualize in Snowflake.

copy github.com/stevenpelley/strava3golang/oauth_client_config.json and the resulting token.json here.

state:
simple demonstration to retrieve the current athlete.

todo:
reorganize into utilities and library code.
create utility to retrieve all activities within a date range, excluding a list of known activities.
create a snowflake account and start defining the schema (dbt?)
utility to ingest activity data into snowflake
simple snowflake dashboard for activities, duration, avg power, bike stress, etc