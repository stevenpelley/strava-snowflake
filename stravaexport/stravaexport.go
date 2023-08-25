package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

const port = 8080 // port of local demo server

// sent a struct{} once we have successfully shutdown the server, at which point
// main may exit
var shutdownChan chan struct{} = make(chan struct{})

func main() {
	initLogging()
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
		log.Panicf("error creating oauth config. %v", err)
	}

	go func() {
		err := initiateOauthRequest(conf)
		if err != nil {
			log.Panicf("error while initiating oauth request: %v", err)
		}
	}()
	// blocks indefinitely

	err = listenAndServeOauthRedirection(listener, conf)
	if err == http.ErrServerClosed {
		// this is the expected case
		slog.Info("server done listening and returned with ErrServerClosed (expected).  " +
			"We will wait for shutdown to return")
		<-shutdownChan
		slog.Info("server shutdown return signalled.  Goodbye.")
	} else {
		slog.Error("unexpected return from server.listen", "error", err)
		os.Exit(1)
	}
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
	slog.Info("confirmed http server started", "response", resp)

	// open a browser and head to strava oauth
	url := conf.AuthCodeURL("state")
	browser.OpenURL(url)
	return nil
}

// Starts a server, using a previously-created Listener, that will handle the redirect
// from Strava oauth
func listenAndServeOauthRedirection(listener net.Listener, conf *oauth2.Config) error {
	// create a server and listen with the previously created listener.  This will
	// server any prior requests that buffered.
	s := &http.Server{}
	s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("confirmed http server started")
		var err error = r.ParseForm()
		if err != nil {
			slog.Error("error parsing form", "error", err)
			w.Write([]byte(fmt.Sprintf("Error parsing the form: %v", err)))
			return
		}

		slog.Info("parsed the form", "form", r.Form, "PostForm", r.PostForm)
		code := r.FormValue("code")
		if code == "" {
			slog.Info("'code' does not exist in form", "Form", r.Form)
			w.Write([]byte(fmt.Sprintf("parsed the form but 'code' does not exist. "+
				"form: %v. PostForm: %v", r.Form, r.PostForm)))
			return
		}
		var token *oauth2.Token
		token, err = conf.Exchange(context.Background(), code)
		if err != nil {
			slog.Error("error exchanging code for token", "error", err, "Form", r.Form)
			w.Write([]byte(fmt.Sprintf("Error exchanging code for token: %v. form: %v",
				err, r.Form)))
			return
		}
		slog.Info("oauth exchange succeeded")
		tokenBytes, err := json.MarshalIndent(token, "", "  ")
		if err != nil {
			slog.Error("error marshalling token", "error", err)
			w.Write([]byte(fmt.Sprintf("Error marshalling token: %v. form: %v",
				err, r.Form)))
			return
		}

		tokenString := string(tokenBytes[:])
		slog.Info("successfully exchanged code for token", "token", tokenString)
		w.Write([]byte(fmt.Sprintf("token: %s", tokenString)))

		// we write this to stdout so that it can be redirected to a file
		// and stored for future use
		file, err := os.OpenFile("token.txt",
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			slog.Error("error opening token.txt file to write token", "error", err)
			w.Write([]byte(fmt.Sprintf("Error opening token.txt to write"+
				"token. error: %v. token: %s", err, tokenString)))
			return
		}
		defer file.Close()
		fmt.Fprint(file, tokenString)

		// we have completed our token exchange and provided/persisted the token.  Shut down
		// the server so that we can end the process.
		go func() {
			err := s.Shutdown(context.Background())
			if err != nil {
				slog.Error("shutting down server failed", "error", err)
				// kill it with fire
				os.Exit(-1)
			}
			slog.Info("successfully shut down the server")
			shutdownChan <- struct{}{}
		}()

		////client := conf.Client(context.Background(), token)
		////stravaApiConfig := stravaapi.NewConfiguration()
		////stravaApiConfig.HTTPClient = client
		////stravaClient := stravaapi.NewAPIClient(stravaApiConfig)
		////detailedAthlete, resp, err := stravaClient.AthletesApi.GetLoggedInAthleteExecute(
		////	stravaClient.AthletesApi.GetLoggedInAthlete(context.Background()))
		////if err != nil {
		////	slog.Error("error retrieving athlete", "error", err)
		////	w.Write([]byte(fmt.Sprintf("Error getting Athlete: %v", err)))
		////	return
		////}
		////w.Write([]byte(fmt.Sprintf("detailed athlete: %+v. resp: %v", detailedAthlete, resp)))
	})

	return s.Serve(listener) // actually start the server, it serves all buffered requests
}
