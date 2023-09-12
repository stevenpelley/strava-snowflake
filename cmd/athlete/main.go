package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/stevenpelley/strava-snowflake/internal/strava"
)

func main() {
	strava.InitLogging("athlete.log")
	stravaClient := strava.CreateStravaClient()
	detailedAthlete, _, err :=
		stravaClient.AthletesAPI.GetLoggedInAthlete(context.Background()).Execute()
	if err != nil {
		log.Panicf("error retrieving athlete: %v", err)
		return
	}

	slog.Info("retrieved athlete", "firstname", *detailedAthlete.Firstname,
		"lastname", *detailedAthlete.Lastname, "id", *detailedAthlete.Id)
}
