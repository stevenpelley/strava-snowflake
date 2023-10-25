package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

func main() {
	strava.InitLogging("activities.log")

	stravaFlags := strava.StravaFlags{}
	err := stravaFlags.InitFlags(flag.CommandLine)
	if err != nil {
		log.Panicf("error initializing flags: %v", err)
	}

	flag.Parse()

	err = stravaFlags.PostProcessFlags(flag.CommandLine)
	if err != nil {
		log.Panicf("error postprocessing flags: %v", err)
	}

	stravaClient, err := strava.CreateStravaClient(&stravaFlags)
	if err != nil {
		log.Panicf("error creating strava client: %v", err)
	}
	stravaFlags.Config.StravaClient = stravaClient

	activities, err := strava.GetActivitiesAndStreams(&stravaFlags.Config, func(activityIds []int64) ([]int64, error) {
		return activityIds, nil
	})
	// may have been a partial result so output the result before checking the error
	for _, aas := range activities {
		fmt.Fprintf(os.Stdout, "%v\n", util.MarshalOrPanic(aas))
	}
	if err != nil {
		log.Panicf("error retrieving activities and streams: %v", err)
	}
}
