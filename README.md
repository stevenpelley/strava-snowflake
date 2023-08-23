# strava-snowflake
track and export personal Strava data.  Ingest and visualize in Snowflake.

Requires openapi-generator (installed from brew on mac)
stravaapi client library generated using `make stravaapi`

state:
created a client that doesn't produce compilation errors.  Nothing else working

need:
means of authenticating with oauth and returning a user's token (and user id?)
a demo retrieving the athletes basic info using the client and oauth