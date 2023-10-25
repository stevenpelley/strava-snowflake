go build ./cmd/activitiessql
# 1 year -8760h
# 2 years -17520h
# 3 years -26280h
#startdurationago="-$((24*(365*3)))h"
startdurationago="-$((24*20))h"
enddurationago="-$((24*(0)))h"
./activitiessql -startdurationago $startdurationago -enddurationago $enddurationago -getstreamsconcurrency 16 -streamstimeoutduration 15m -activitiestimeoutduration 5m snowflake -configfilename snowflake_config.json -etltablename '"STRAVA"."ACTIVITIES"."ETL"'
