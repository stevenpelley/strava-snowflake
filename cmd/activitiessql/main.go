package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stevenpelley/strava-snowflake/internal/intsql/intduckdb"
	"github.com/stevenpelley/strava-snowflake/internal/intsql/intsnowflake"
	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

func main() {
	strava.InitLogging("activitiessql.log")

	// first parse global options.  This contains global flags.  Then flag.Args()[0]
	// contains the subcommand.  flag.Args()[1:] contains "args" to be passed to
	// FlagSet.Parse().
	stravaFlags := strava.StravaFlags{}
	err := stravaFlags.InitFlags(flag.CommandLine)
	if err != nil {
		log.Panicf("error preprocessing strava flags: %v", err)
	}
	var mergeOnly bool
	flag.BoolVar(&mergeOnly, "mergeonly", false, "if true perform merge from existing data.  Do not retrieve.")
	flag.Parse()

	err = stravaFlags.PostProcessFlags(flag.CommandLine)
	if err != nil {
		log.Panicf("error postprocessing strava flags: %v", err)
	}

	config := &stravaFlags.Config

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Please specify a subcommand.")
	}
	cmd, args := args[0], args[1:]
	sdb, err := subcommand(cmd, args)
	if err != nil {
		log.Panicf("creating subcommand StravaDatabase: %v", err)
	}

	err = sdb.OpenDB()
	if err != nil {
		log.Panicf("StravaDatabase.OpenDB: %v", err)
	}

	defer sdb.Close()

	var getActivitiesErr error
	if !mergeOnly {
		err = sdb.InitAndValidateSchema()
		if err != nil {
			log.Panicf("error initializing data schema: %v", err)
		}

		stravaClient, err := strava.CreateStravaClient(&stravaFlags)
		if err != nil {
			log.Panicf("error creating strava client: %v", err)
		}
		config.StravaClient = stravaClient

		// recall that this may return a partial result alongside an error

		var activities []util.Jsonable
		activities, getActivitiesErr = strava.GetActivitiesAndStreams(config, sdb.FilterKnownActivityIds)

		if len(activities) == 0 {
			if getActivitiesErr != nil {
				log.Panicf("GetActivitiesAndStreams: %v", getActivitiesErr)
			}
			fmt.Println("no new activities found")
			return
		}

		err = sdb.UploadActivityJson(activities)
		if err != nil {
			log.Panicf("error loading json: %v", err)
		}
		fmt.Printf("uploaded %v activities\n", len(activities))
	}

	err = sdb.MergeActivities()
	if err != nil {
		log.Panicf("error merging activities: %v", err)
	}

	// note any error from acquiring the activities earlier
	if getActivitiesErr != nil {
		log.Panicf("GetActivitiesAndStreams, attempted partial result: %v", getActivitiesErr)
	}
}

func subcommand(cmd string, args []string) (intsql.StravaDatabase, error) {
	flagSet := flag.NewFlagSet(cmd, flag.ExitOnError)
	var f intsql.SqlFlags
	switch cmd {
	case "duckdb":
		f = &intduckdb.DuckdbFlags{}
	case "snowflake":
		f = &intsnowflake.SFFlags{}
	}

	f.InitFlags(flagSet)
	flagSet.Parse(args)
	f.PostProcessFlags(flagSet)
	return f.NewStravaDatabase()
}
