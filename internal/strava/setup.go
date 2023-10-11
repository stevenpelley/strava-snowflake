package strava

import (
	"context"
	"encoding/json"
	"io"
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
	var writer io.Writer = nil
	if logFile == "" {
		// use stderr
		writer = os.Stderr
	} else {
		// If the file doesn't exist, create it or append to the file
		file, err := os.OpenFile(logFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		writer = file
	}

	handlerOptions := slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	logger := slog.New(slog.NewJSONHandler(writer, &handlerOptions))
	slog.SetDefault(logger)
}
