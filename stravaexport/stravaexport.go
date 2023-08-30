package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/stevenpelley/strava3golang"
	"golang.org/x/oauth2"
)

func main() {
	initLogging()
	conf, err := createOauthConfig()
	if err != nil {
		log.Panicf("panic while creating oauth config: %v", err)
	}

	token, err := createToken()
	if err != nil {
		log.Panicf("panic while creating token: %v", err)
	}

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

func initLogging() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("stravaexport.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))
	slog.SetDefault(logger)
}

// Create the oauth config used for all oauth requests
func createOauthConfig() (*oauth2.Config, error) {
	var err error
	data, err := os.ReadFile("oauth_client_config.json")
	if err != nil {
		log.Panicf("error reading the oauth client config: %v", err)
	}

	var clientConfig oauth2.Config
	err = json.Unmarshal(data, &clientConfig)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %w", err)
	}

	return &clientConfig, nil
}

func createToken() (*oauth2.Token, error) {
	var err error
	data, err := os.ReadFile("token.json")
	if err != nil {
		log.Panicf("error reading the token: %v", err)
	}

	var token oauth2.Token
	err = json.Unmarshal(data, &token)
	if err != nil {
		return nil, fmt.Errorf("error parsing token json: %w", err)
	}

	return &token, nil
}
