package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava3golang"
)

func main() {
	strava.InitLogging("strava.log")
	conf := strava.CreateOauthConfig()
	token := strava.CreateToken()

	client := conf.Client(context.Background(), token)
	stravaApiConfig := strava3golang.NewConfiguration()
	stravaApiConfig.HTTPClient = client
	stravaClient := strava3golang.NewAPIClient(stravaApiConfig)
	detailedAthlete, _, err := stravaClient.AthletesAPI.GetLoggedInAthleteExecute(
		stravaClient.AthletesAPI.GetLoggedInAthlete(context.Background()))
	if err != nil {
		log.Panicf("error retrieving athlete: %v", err)
		return
	}

	slog.Info("retrieved athlete", "firstname", *detailedAthlete.Firstname,
		"lastname", *detailedAthlete.Lastname, "id", *detailedAthlete.Id)
}
