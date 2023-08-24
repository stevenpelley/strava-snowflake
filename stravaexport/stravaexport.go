package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

const port = 8080 // port of local demo server

func main() {
	var err error
	// Try to bind address.  This allows subsequent requests, for example our own check that the
	// server has started, to buffer.
	// we reuse the listener later when creating the server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	conf, err := createOauthConfig()
	if err != nil {
		panic(fmt.Sprintf("error creating oauth config. %v", err))
	}

	go func() {
		err := initiateOauthRequest(conf)
		if err != nil {
			panic(fmt.Sprintf("error while initiating oauth request: %v", err))
		}
	}()
	// blocks indefinitely
	log.Fatal(listenAndServeOauthRedirection(listener))
}

// Create the oauth config used for all oauth requests
func createOauthConfig() (*oauth2.Config, error) {
	var err error
	data, err := os.ReadFile("oauth_client_config.json")
	if err != nil {
		panic(fmt.Sprintf("error reading the oauth client config: %v", err))
	}
	type ClientConfig struct {
		ClientId     string
		ClientSecret string
	}
	var clientConfig ClientConfig
	err = json.Unmarshal(data, &clientConfig)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %w", err)
	}

	if clientConfig.ClientId == "" {
		return nil, fmt.Errorf("oauth client config ClientId is empty. data: %v", data)
	}

	if clientConfig.ClientSecret == "" {
		return nil, fmt.Errorf("oauth client config ClientSecret is empty. data: %v", data)
	}

	return &oauth2.Config{
		ClientID:     clientConfig.ClientId,
		ClientSecret: clientConfig.ClientSecret,
		// Strava oauth does not conform to the oauth2 spec.
		// It expects comma-delimited instead of space-delimited scopes.
		// Putting all scopes in a single string is a workaround
		Scopes:      []string{"activity:read_all"},
		Endpoint:    endpoints.Strava,
		RedirectURL: fmt.Sprintf("http://localhost:%d", port),
	}, nil
}

// Starts the oauth process by opening a link in a browser directing the user to Strava's oauth
func initiateOauthRequest(conf *oauth2.Config) error {
	// block until we determine that we are properly serving http so that we can handle the
	// oauth redirect
	var err error
	var duration time.Duration
	duration, err = time.ParseDuration("10s")
	if err != nil {
		return fmt.Errorf("error parsing http client duration: %w", err)
	}
	client := &http.Client{
		Timeout: duration,
	}

	resp, err := client.Get(fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		return fmt.Errorf("error waiting for http server to start: %w", err)
	}
	fmt.Println("http server started. resp", resp)

	// open a browser and head to strava oauth
	url := conf.AuthCodeURL("state")
	browser.OpenURL(url)
	return nil
}

// Starts a server, using a previously-created Listener, that will handle the redirect
// from Strava oauth
func listenAndServeOauthRedirection(listener net.Listener) error {
	// create a server and listen with the previously created listener.  This will
	// server any prior requests that buffered.
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("incoming request", r)
			var err error = r.ParseForm()
			if err != nil {
				fmt.Println("error parsing form", err, "request", r)
				w.Write([]byte(fmt.Sprintf("Error parsing the form: %v", err)))
			} else {
				fmt.Println("parsed the form.", "form", r.Form, "PostForm", r.PostForm)
				w.Write([]byte(fmt.Sprintf("parsed the form. form: %v. PostForm: %v", r.Form, r.PostForm)))
			}
		}),
	}

	return s.Serve(listener) // actually start the server, it manges all buffered requests
}
