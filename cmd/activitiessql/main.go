package main

import (
	"flag"
	"log"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stevenpelley/strava-snowflake/internal/strava"
)

func main() {
	strava.InitLogging("activitiessql.log")

	stravaFlags := strava.StravaFlags{}
	stravaFlags.InitFlags()
	duckdbFlags := intsql.DuckdbFlags{}
	duckdbFlags.InitFlags()

	flag.Parse()

	stravaFlags.PostProcessFlags()
	config := &stravaFlags.Config
	duckdbFlags.PostProcessFlags()

	db, err := intsql.OpenDB(duckdbFlags.DbFileName)
	if err != nil {
		log.Panicf("error opening db file %v: %v", duckdbFlags.DbFileName, err)
	}
	defer db.Close()

	err = intsql.InitAndValidateSchema(db)
	if err != nil {
		log.Panicf("error initializing data schema: %v", err)
	}

	activityIdsToIgnore, err := intsql.GetExistingActivityIds(db)
	if err != nil {
		log.Panicf("error loading existing activity ids: %v", err)
	}
	for activityId, _ := range activityIdsToIgnore {
		config.ActivityIdsToIgnore[activityId] = struct{}{}
	}

	stravaClient, err := strava.CreateStravaClient(&stravaFlags)
	if err != nil {
		log.Panicf("error creating strava client: %v", err)
	}
	config.StravaClient = stravaClient

	// recall that this may return a partial result alongside an error
	activities, getActivitiesErr := strava.GetActivitiesAndStreams(config)

	err = intsql.UploadActivityJson(db, activities)
	if err != nil {
		log.Panicf("error loading json: %v", err)
	}

	err = intsql.MergeActivities(db, duckdbFlags.StreamsEtlLimit)
	if err != nil {
		log.Panicf("error merging activities: %v", err)
	}

	// note any error from acquiring the activities earlier
	if getActivitiesErr != nil {
		log.Panicf("error retrieving activities and streams: %v", getActivitiesErr)
	}
}
