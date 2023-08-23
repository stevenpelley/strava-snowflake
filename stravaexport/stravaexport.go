package main

import (
	"fmt"
)

//const port = 8080 // port of local demo server

func main() {
	fmt.Println("test")

	// test/demo to redirect to a browser
	//browser.OpenURL("https://www.google.com")

	// test/demo taken from the stravaapi generated README, but I think this is pseudocode.
	////var token oauth2.Token
	////tokenSource := oauth2.Config.TokenSource(oauth2.createContext(httpClient), &token)
	////auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
	////r, err := client.Service.Operation(auth, args)

	// test/demo from strava API v2.  Gives the general workflow but doesn't work.
	////	// setup the credentials for your app
	////	// These need to be set to reflect your application
	////	// and can be found at https://www.strava.com/settings/api
	////	flag.IntVar(&strava.ClientId, "id", 0, "Strava Client ID")
	////	flag.StringVar(&strava.ClientSecret, "secret", "", "Strava Client Secret")
	////
	////	flag.Parse()
	////
	////	if strava.ClientId == 0 || strava.ClientSecret == "" {
	////		fmt.Println("\nPlease provide your application's client_id and client_secret.")
	////		fmt.Println("For example: go run oauth_example.go -id=9 -secret=longrandomsecret")
	////		fmt.Println(" ")
	////
	////		flag.PrintDefaults()
	////		os.Exit(1)
	////	}
	////
	////// define a strava.OAuthAuthenticator to hold state.
	////// The callback url is used to generate an AuthorizationURL.
	////// The RequestClientGenerator can be used to generate an http.RequestClient.
	////// This is usually when running on the Google App Engine platform.
	////authenticator = &strava.OAuthAuthenticator{
	////	CallbackURL:            fmt.Sprintf("http://localhost:%d/exchange_token", port),
	////	RequestClientGenerator: nil,
	////}

	////http.HandleFunc("/", indexHandler)

	////path, err := authenticator.CallbackPath()
	////if err != nil {
	////	// possibly that the callback url set above is invalid
	////	fmt.Println(err)
	////	os.Exit(1)
	////}
	////http.HandleFunc(path, authenticator.HandlerFunc(oAuthSuccess, oAuthFailure))

	////// start the server
	////fmt.Printf("Visit http://localhost:%d/ to view the demo\n", port)
	////fmt.Printf("ctrl-c to exit")
	////http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}

////func indexHandler(w http.ResponseWriter, r *http.Request) {
////	// you should make this a template in your real application
////	fmt.Fprintf(w, `<a href="%s">`, authenticator.AuthorizationURL("state1", strava.Permissions.Public, true))
////	fmt.Fprint(w, `<img src="http://strava.github.io/api/images/ConnectWithStrava.png" />`)
////	fmt.Fprint(w, `</a>`)
////}
////
////func oAuthSuccess(auth *strava.AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
////	fmt.Fprintf(w, "SUCCESS:\nAt this point you can use this information to create a new user or link the account to one of your existing users\n")
////	fmt.Fprintf(w, "State: %s\n\n", auth.State)
////	fmt.Fprintf(w, "Access Token: %s\n\n", auth.AccessToken)
////
////	fmt.Fprintf(w, "The Authenticated Athlete (you):\n")
////	content, _ := json.MarshalIndent(auth.Athlete, "", " ")
////	fmt.Fprint(w, string(content))
////}
////
////func oAuthFailure(err error, w http.ResponseWriter, r *http.Request) {
////	fmt.Fprintf(w, "Authorization Failure:\n")
////
////	// some standard error checking
////	if err == strava.OAuthAuthorizationDeniedErr {
////		fmt.Fprint(w, "The user clicked the 'Do not Authorize' button on the previous page.\n")
////		fmt.Fprint(w, "This is the main error your application should handle.")
////	} else if err == strava.OAuthInvalidCredentialsErr {
////		fmt.Fprint(w, "You provided an incorrect client_id or client_secret.\nDid you remember to set them at the begininng of this file?")
////	} else if err == strava.OAuthInvalidCodeErr {
////		fmt.Fprint(w, "The temporary token was not recognized, this shouldn't happen normally")
////	} else if err == strava.OAuthServerErr {
////		fmt.Fprint(w, "There was some sort of server error, try again to see if the problem continues")
////	} else {
////		fmt.Fprint(w, err)
////	}
////}
