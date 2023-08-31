package strava

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"os"

	"github.com/stevenpelley/strava3golang"
	"golang.org/x/oauth2"
)

// Create the oauth config used for all oauth requests
func CreateOauthConfig() *oauth2.Config {
	var err error
	data, err := os.ReadFile("oauth_client_config.json")
	if err != nil {
		log.Panicf("error reading the oauth client config: %v", err)
	}

	var clientConfig oauth2.Config
	err = json.Unmarshal(data, &clientConfig)
	if err != nil {
		log.Panicf("panic while parsing oauth config: %v", err)
	}

	return &clientConfig
}

func CreateToken() *oauth2.Token {
	var err error
	data, err := os.ReadFile("token.json")
	if err != nil {
		log.Panicf("error reading the token: %v", err)
	}

	var token oauth2.Token
	err = json.Unmarshal(data, &token)
	if err != nil {
		log.Panicf("error parsing the token: %v", err)
	}

	return &token
}

func CreateStravaClient() *strava3golang.APIClient {
	conf := CreateOauthConfig()
	token := CreateToken()

	client := conf.Client(context.Background(), token)
	stravaApiConfig := strava3golang.NewConfiguration()
	stravaApiConfig.HTTPClient = client
	return strava3golang.NewAPIClient(stravaApiConfig)
}

func InitLogging(logFile string) {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))
	slog.SetDefault(logger)
}
