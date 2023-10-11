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
	stravaFlags.InitFlags()
	flag.Parse()
	stravaFlags.PostProcessFlags()

	stravaClient := strava.CreateStravaClient()
	stravaFlags.Config.StravaClient = stravaClient

	activities, err := strava.GetActivitiesAndStreams(&stravaFlags.Config)
	// may have been a partial result so output the result before checking the error
	for _, aas := range activities {
		fmt.Fprintf(os.Stdout, "%v\n", util.MarshalOrPanic(aas))
	}
	if err != nil {
		log.Panicf("error retrieving activities and streams: %v", err)
	}
}
