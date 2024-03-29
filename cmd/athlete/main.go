package main

import (
	"context"
	"flag"
	"log"
	"log/slog"

	"github.com/stevenpelley/strava-snowflake/internal/strava"
)

func main() {
	strava.InitLogging("athlete.log")
	// most of these are ignored but this is just a simple demo
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
	detailedAthlete, _, err :=
		stravaClient.AthletesAPI.GetLoggedInAthlete(context.Background()).Execute()
	if err != nil {
		log.Panicf("error retrieving athlete: %v", err)
		return
	}

	slog.Info("retrieved athlete", "firstname", *detailedAthlete.Firstname,
		"lastname", *detailedAthlete.Lastname, "id", *detailedAthlete.Id)
}
