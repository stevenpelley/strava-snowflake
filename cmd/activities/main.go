package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

var activityIdsToIgnoreSlice = make([]int64, 0)
var startDurationAgo time.Duration
var endDurationAgo time.Duration
var ignoreIdsFile string

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

	flag.StringVar(&ignoreIdsFile, "ignoreidsfile", "", "file containing newline delimited activity ids to ignore.  Recommend creating this file with jq .Activities.id from previous runs")

	flag.Parse()

	if startDurationAgo > 0 {
		log.Panicf("startDurationAgo must be nonpositive (cannot be in the future): %v",
			startDurationAgo)
	}
	if endDurationAgo > 0 {
		log.Panicf("startDurationAgo must be nonpositive (cannot be in the future): %v",
			endDurationAgo)
	}
	if startDurationAgo > endDurationAgo {
		log.Panicf("startDurationAgo is more recent than endDurationAgo. start: %v. end: %v",
			startDurationAgo, endDurationAgo)
	}

	t := time.Now()
	config.StartTime = t.Add(startDurationAgo)
	config.EndTime = t.Add(endDurationAgo)

	if ignoreIdsFile != "" {
		f, err := os.Open(ignoreIdsFile)
		if err != nil {
			log.Panicf("failed to open file %v: %v", ignoreIdsFile, err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			s := scanner.Text()
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Panicf("failed to convert line to int.  line: %v. error: %v", s, err)
			}
			activityIdsToIgnoreSlice = append(activityIdsToIgnoreSlice, int64(i))
		}
		if err = scanner.Err(); err != nil {
			log.Panicf("error while scanning ignoreidsfile: %v", err)
		}
	}

	config.ActivityIdsToIgnore = make(map[int64]struct{})
	for _, i := range activityIdsToIgnoreSlice {
		config.ActivityIdsToIgnore[i] = struct{}{}
	}
	slog.Debug("initialized inputs", "slice", activityIdsToIgnoreSlice, "config", config)
}

func main() {
	strava.InitLogging("activities.log")
	initInputs()
	stravaClient := strava.CreateStravaClient()
	config.StravaClient = stravaClient

	activities, err := strava.GetActivitiesAndStreams(&config)
	// may have been a partial result so output the result before checking the error
	for _, aas := range activities {
		fmt.Fprintf(os.Stdout, "%v\n", util.MarshalOrPanic(aas))
	}
	if err != nil {
		log.Panicf("error retrieving activities and streams: %v", err)
	}

}
