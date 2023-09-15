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

var startDurationAgo time.Duration
var endDurationAgo time.Duration
var activityIdsToIgnore = map[int64]struct{}{}

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

	flag.Func("ignoreid", "ignore activity id (may be used multiple times)", func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		activityIdsToIgnore[int64(i)] = struct{}{}
		return nil
	})

	flag.Parse()
}

func main() {
	strava.InitLogging("activities.log")
	initInputs()
	stravaClient := strava.CreateStravaClient()
	t := time.Now()
	startTime := t.Add(startDurationAgo)
	endTime := t.Add(endDurationAgo)

	activities, err := strava.GetActivitiesAndStreams(
		stravaClient, startTime, endTime, activityIdsToIgnore)

	if err != nil {
		log.Panicf("error retrieving activities and streams: %v", err)
	}

	for _, aas := range activities {
		fmt.Fprintf(os.Stdout, "%v\n", util.MarshalOrPanic(aas))
	}
}
