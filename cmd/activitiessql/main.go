package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/stevenpelley/strava-snowflake/internal/intsql/intduckdb"
	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

func main() {
	strava.InitLogging("activitiessql.log")

	stravaFlags := strava.StravaFlags{}
	duckdbFlags := intduckdb.DuckdbFlags{}
	err := util.InitAllFlags(&stravaFlags, &duckdbFlags)
	if err != nil {
		log.Panicf("error initializing flags: %v", err)
	}

	var mergeOnly bool
	flag.BoolVar(&mergeOnly, "mergeonly", false, "if true perform merge from existing data.  Do not retrieve.")

	flag.Parse()

	err = util.PostProcessAllFlags(&stravaFlags, &duckdbFlags)
	if err != nil {
		log.Panicf("error postprocessing flags: %v", err)
	}

	config := &stravaFlags.Config

	var sdb intduckdb.DuckdbStrava
	err = sdb.OpenDB(duckdbFlags.DbFileName)
	if err != nil {
		log.Panicf("error opening db file %v: %v", duckdbFlags.DbFileName, err)
	}
	defer sdb.Close()

	var getActivitiesErr error
	if !mergeOnly {
		err = sdb.InitAndValidateSchema()
		if err != nil {
			log.Panicf("error initializing data schema: %v", err)
		}

		stravaClient, err := strava.CreateStravaClient(&stravaFlags)
		if err != nil {
			log.Panicf("error creating strava client: %v", err)
		}
		config.StravaClient = stravaClient

		// recall that this may return a partial result alongside an error

		var activities []util.Jsonable
		activities, getActivitiesErr = strava.GetActivitiesAndStreams(config, sdb.FilterKnownActivityIds)

		if len(activities) == 0 {
			if getActivitiesErr != nil {
				log.Panicf(": %v", getActivitiesErr)
			}
			fmt.Println("no new activities found")
			return
		}

		err = sdb.UploadActivityJson(activities)
		if err != nil {
			log.Panicf("error loading json: %v", err)
		}
	}

	err = sdb.MergeActivities()
	if err != nil {
		log.Panicf("error merging activities: %v", err)
	}

	// note any error from acquiring the activities earlier
	if getActivitiesErr != nil {
		log.Panicf("error retrieving activities and streams: %v", getActivitiesErr)
	}
}
