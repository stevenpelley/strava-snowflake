# strava-snowflake
track and export personal Strava data.  Ingest and visualize in Snowflake.

Requires openapi-generator (installed from brew on mac)
stravaapi client library generated using `make stravaapi`

copy stravaexport/oauth_client_config.json.template into strava/export/oauth_client_config.json and place your Strava app ClientID and ClientSecret as the corresponding values.  This non-template file is in .gitignore so that secrets stay out of the repository.

state:
Created a client that doesn't produce compilation errors.
stravaexport performs oauth initiation and redirect.  Manually verified that I get a code in the redirect request.

need:
use the provided oauth code to get a refresh token, access token, and user info.
a demo retrieving the athletes basic info using the client and oauth
refactor to put the strava api in a separate directory.  This was painful enough to create that others may find it useful.
create utility to retrieve all activities within a date range, excluding a list of known activities.