clean_strava:
	rm -rf strava

strava:
	# broken -- enum values collide
	#swagger-codegen generate --input-spec https://developers.strava.com/swagger/swagger.json --lang go --output strava
	# works!
	openapi-generator generate --input-spec https://developers.strava.com/swagger/swagger.json --generator-name go -p enumClassPrefix=true -p isGoSubmodule=true --output stravaapi --skip-validate-spec --package-name stravaapi --git-host github.com --git-user-id stevenpelley --git-repo-id snowflake-strava 

# pursue in the future, switch to openapi
strava_from_docker:
	./swagger.sh generate client -f "https://developers.strava.com/swagger/swagger.json"