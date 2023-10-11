package strava

import (
	"bufio"
	"flag"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type StravaFlags struct {
	activityIdsToIgnoreSlice []int64
	startDurationAgo         time.Duration
	endDurationAgo           time.Duration
	ignoreIdsFile            string
	Config                   ActivitiesConfig
}

// must be called before flag.Parse().  Panics on any error
func (sf *StravaFlags) InitFlags() {
	defaultStartDurationAgo, err := time.ParseDuration("-168h")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(&sf.startDurationAgo, "startdurationago", defaultStartDurationAgo,
		"duration relative to \"now\" to start retrieving activities")

	defaultEndDurationAgo, err := time.ParseDuration("0")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(&sf.endDurationAgo, "enddurationago", defaultEndDurationAgo,
		"duration relative to \"now\" to stop retrieving activities")

	flag.IntVar(&sf.Config.GetStreamsConcurrency, "getstreamsconcurrency", 4,
		"maximum concurrency with which to retrieve activity streams")

	defaultActivitiesTimeoutDuration, err := time.ParseDuration("30s")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(
		&sf.Config.ActivitiesTimeoutDuration,
		"activitiestimeoutduration",
		defaultActivitiesTimeoutDuration,
		"duration before timing out to retrieve activities")

	defaultStreamsTimeoutDuration, err := time.ParseDuration("60s")
	if err != nil {
		log.Panicf("error construction flag default duration: %v", err)
	}
	flag.DurationVar(
		&sf.Config.StreamsTimeoutDuration,
		"streamstimeoutduration",
		defaultStreamsTimeoutDuration,
		"duration before timing out to retrieve streams")

	flag.IntVar(&sf.Config.ActivitiesPerPage, "activitiesperpage", 30, "activities per page used in http requests")

	flag.DurationVar(&sf.Config.PreStreamSleep, "prestreamsleep", 0,
		"duration to sleep prior to retrieving each stream to act as throttling")

	flag.Func("ignoreid", "ignore activity id (may be used multiple times)", func(s string) error {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		sf.activityIdsToIgnoreSlice = append(sf.activityIdsToIgnoreSlice, int64(i))
		return nil
	})

	flag.StringVar(&sf.ignoreIdsFile, "ignoreidsfile", "", "file containing newline delimited activity ids to ignore.  Recommend creating this file with jq .Activities.id from previous runs")
}

// must be called after flag.Parse().  Panics on any error
func (sf *StravaFlags) PostProcessFlags() {
	if sf.startDurationAgo > 0 {
		log.Panicf("startDurationAgo must be nonpositive (cannot be in the future): %v",
			sf.startDurationAgo)
	}
	if sf.endDurationAgo > 0 {
		log.Panicf("startDurationAgo must be nonpositive (cannot be in the future): %v",
			sf.endDurationAgo)
	}
	if sf.startDurationAgo > sf.endDurationAgo {
		log.Panicf("startDurationAgo is more recent than endDurationAgo. start: %v. end: %v",
			sf.startDurationAgo, sf.endDurationAgo)
	}

	t := time.Now()
	sf.Config.StartTime = t.Add(sf.startDurationAgo)
	sf.Config.EndTime = t.Add(sf.endDurationAgo)

	if sf.ignoreIdsFile != "" {
		f, err := os.Open(sf.ignoreIdsFile)
		if err != nil {
			log.Panicf("failed to open file %v: %v", sf.ignoreIdsFile, err)
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			s := scanner.Text()
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Panicf("failed to convert line to int.  line: %v. error: %v", s, err)
			}
			sf.activityIdsToIgnoreSlice = append(sf.activityIdsToIgnoreSlice, int64(i))
		}
		if err = scanner.Err(); err != nil {
			log.Panicf("error while scanning ignoreidsfile: %v", err)
		}
	}

	sf.Config.ActivityIdsToIgnore = make(IntSet)
	for _, i := range sf.activityIdsToIgnoreSlice {
		sf.Config.ActivityIdsToIgnore[i] = struct{}{}
	}
	slog.Debug("initialized inputs", "slice", sf.activityIdsToIgnoreSlice, "config", sf.Config)
}
