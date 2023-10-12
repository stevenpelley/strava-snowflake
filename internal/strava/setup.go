package strava

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/stevenpelley/strava3golang"
	"golang.org/x/oauth2"
)

// load the file, expected to be json-encoded data.  Unmarshal into object pointed to by destination
func readAndUnmarshalConfigFile(fileName string, destination any) error {
	if fileName == "" {
		return errors.New("readAndUnmarshalConfigFile requires non empty oauth config file name")
	}
	var err error
	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("reading config. file: %v: %w", fileName, err)
	}
	err = json.Unmarshal(data, destination)
	if err != nil {
		return fmt.Errorf("parsing config. file: %v: %w", fileName, err)
	}
	return nil
}

// Create the oauth config used for all oauth requests
func CreateOauthConfig(oauthConfigFile string) (*oauth2.Config, error) {
	var clientConfig oauth2.Config
	err := readAndUnmarshalConfigFile(oauthConfigFile, &clientConfig)
	if err != nil {
		return nil, fmt.Errorf("CreateOauthConfig: %w", err)
	}
	return &clientConfig, nil
}

func CreateToken(tokenFile string) (*oauth2.Token, error) {
	var token oauth2.Token
	err := readAndUnmarshalConfigFile(tokenFile, &token)
	if err != nil {
		return nil, fmt.Errorf("CreateToken: %w", err)
	}
	return &token, nil
}

func CreateStravaClient(sf *StravaFlags) (*strava3golang.APIClient, error) {
	conf, err := CreateOauthConfig(sf.OauthConfigFile)
	if err != nil {
		return nil, fmt.Errorf("CreateStravaClient: %w", err)
	}
	token, err := CreateToken(sf.OauthTokenFile)
	if err != nil {
		return nil, fmt.Errorf("CreateStravaClient: %w", err)
	}

	client := conf.Client(context.Background(), token)
	stravaApiConfig := strava3golang.NewConfiguration()
	stravaApiConfig.HTTPClient = client
	return strava3golang.NewAPIClient(stravaApiConfig), nil
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
