package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

var activityIdsToIgnoreSlice = make([]int64, 0)
var startDurationAgo time.Duration
var endDurationAgo time.Duration

var config strava.ActivitiesConfig = strava.ActivitiesConfig{}

func initInputs() {
	defaultStartDurationAgo, err := time.ParseDuration("-168h")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(&startDurationAgo, "startdurationago", defaultStartDurationAgo,
		"duration relative to \"now\" to start retrieving activities")

	defaultEndDurationAgo, err := time.ParseDuration("0")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(&endDurationAgo, "enddurationago", defaultEndDurationAgo,
		"duration relative to \"now\" to stop retrieving activities")

	flag.IntVar(&config.GetStreamsConcurrency, "getstreamsconcurrency", 4,
		"maximum concurrency with which to retrieve activity streams")

	defaultActivitiesTimeoutDuration, err := time.ParseDuration("30s")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(
		&config.ActivitiesTimeoutDuration,
		"activitiestimeoutduration",
		defaultActivitiesTimeoutDuration,
		"duration before timing out to retrieve activities")

	defaultStreamsTimeoutDuration, err := time.ParseDuration("60s")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(
		&config.StreamsTimeoutDuration,
		"streamstimeoutduration",
		defaultStreamsTimeoutDuration,
		"duration before timing out to retrieve streams")

	flag.IntVar(&config.ActivitiesPerPage, "activitiesperpage", 30, "activities per page used in http requests")

	flag.DurationVar(&config.PreStreamSleep, "prestreamsleep", 0,
		"duration to sleep prior to retrieving each stream to act as throttling")

	flag.Func("ignoreid", "ignore activity id (may be used multiple times)", func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		activityIdsToIgnoreSlice = append(activityIdsToIgnoreSlice, int64(i))
		return nil
	})

	flag.Parse()

	t := time.Now()
	config.StartTime = t.Add(startDurationAgo)
	config.EndTime = t.Add(endDurationAgo)

	for i, _ := range activityIdsToIgnoreSlice {
		config.ActivityIdsToIgnore[int64(i)] = struct{}{}
	}
}

func main() {
	strava.InitLogging("activities.log")
	initInputs()
	stravaClient := strava.CreateStravaClient()
	config.StravaClient = stravaClient

	activities, err := strava.GetActivitiesAndStreams(&config)

	if err != nil {
		log.Panicf("error retrieving activities and streams: %v", err)
	}

	for _, aas := range activities {
		fmt.Fprintf(os.Stdout, "%v\n", util.MarshalOrPanic(aas))
	}
}
