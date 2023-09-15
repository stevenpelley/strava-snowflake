package strava

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"slices"
	"time"

	"github.com/fatih/structs"
	"github.com/stevenpelley/strava-snowflake/internal/util"
	"github.com/stevenpelley/strava3golang"
	"golang.org/x/sync/errgroup"
)

const activitiesPageSize int32 = 30

func RetrieveActivities(stravaClient *strava3golang.APIClient, start time.Time, end time.Time) (
	[]strava3golang.SummaryActivity, error) {
	startUnixTime := start.UTC().Unix()
	endUnixTime := end.UTC().Unix()
	after := int32(startUnixTime)
	before := int32(endUnixTime)

	var currentPage int32 = 1
	var itemsInPage int = 0
	allActivities := make([]strava3golang.SummaryActivity, 0)
	timeoutDuration, err := time.ParseDuration("30s")
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancelFunc()

	for currentPage == 1 || int32(itemsInPage) == activitiesPageSize {
		activities, _, err := stravaClient.ActivitiesAPI.GetLoggedInAthleteActivities(
			ctx).
			After(after).
			Before(before).
			PerPage(activitiesPageSize).
			Page(currentPage).
			Execute()
		if err != nil {
			return nil, err
		}
		allActivities = append(allActivities, activities...)
		itemsInPage = len(activities)
		currentPage++
	}

	return allActivities, nil
}

func RetrieveStreams(stravaClient *strava3golang.APIClient, activityIds []int64) (
	[]*strava3golang.StreamSet, error) {
	d, err := time.ParseDuration("10s")
	if err != nil {
		log.Panicf("error parsing duration: %v", err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), d)
	defer cancelFunc()
	g, ctx := errgroup.WithContext(ctx)
	streams := make([]*strava3golang.StreamSet, len(activityIds))
	for i, activityId := range activityIds {
		i, activityId := i, activityId // bind to iteration-specific variable for closure
		g.Go(func() error {
			streamSet, _, err := stravaClient.StreamsAPI.GetActivityStreams(ctx, activityId).
				KeyByType(true).
				Keys(AllStreamKindsString).
				Execute()
			if err != nil {
				return fmt.Errorf(
					"error retrieving activity stream. id: %v. activity index: %v. error: %w",
					activityId, i, err)
			}
			streams[i] = streamSet
			return nil
		})
	}
	err = g.Wait()
	return streams, err
}

// intended to be used for logging activities.
type AbridgedSummaryActivity struct {
	// The unique identifier of the activity
	Id *int64 `json:"id,omitempty"`
	// The identifier provided at upload time
	ExternalId *string `json:"external_id,omitempty"`
	// The name of the activity
	Name *string `json:"name,omitempty"`
	// The activity's distance, in meters
	Distance *float32 `json:"distance,omitempty"`
	// The activity's moving time, in seconds
	MovingTime *int32 `json:"moving_time,omitempty"`
	// The activity's elapsed time, in seconds
	ElapsedTime *int32 `json:"elapsed_time,omitempty"`
	// The activity's average speed, in meters per second
	AverageSpeed *float32 `json:"average_speed,omitempty"`
	// Average power output in watts during this activity. Rides only
	AverageWatts *float32                    `json:"average_watts,omitempty"`
	Type         *strava3golang.ActivityType `json:"type,omitempty"`
}

func GetAbridgedActivity(o *strava3golang.SummaryActivity) AbridgedSummaryActivity {
	return AbridgedSummaryActivity{
		Id:           o.Id,
		ExternalId:   o.ExternalId,
		Name:         o.Name,
		Distance:     o.Distance,
		MovingTime:   o.MovingTime,
		ElapsedTime:  o.ElapsedTime,
		AverageSpeed: o.AverageSpeed,
		AverageWatts: o.AverageWatts,
		Type:         o.Type,
	}
}

func MapSummaryActivitiesToAbridged(
	activities []strava3golang.SummaryActivity) []AbridgedSummaryActivity {
	vals := make([]AbridgedSummaryActivity, len(activities))
	for i, a := range activities {
		vals[i] = GetAbridgedActivity(&a)
	}
	return vals
}

type ActivityAndStream struct {
	Activity  *strava3golang.SummaryActivity
	StreamSet *strava3golang.StreamSet
}

func GetActivitiesAndStreams(
	stravaClient *strava3golang.APIClient,
	startTime time.Time,
	endTime time.Time,
	activityIdsToIgnore map[int64]struct{}) (
	[]ActivityAndStream, error) {

	activities, err := RetrieveActivities(stravaClient, startTime, endTime)
	if err != nil {
		return nil, err
	}

	slog.Info("retrieved activities",
		"length", len(activities),
		"activities", util.LogValueFunc(func() any {
			return MapSummaryActivitiesToAbridged(activities)
		}))

	// filter out activities to be ignored
	activities = slices.DeleteFunc(activities, func(a strava3golang.SummaryActivity) bool {
		_, ok := activityIdsToIgnore[*a.Id]
		return ok
	})

	slog.Info("filtered activities",
		"length", len(activities),
		"activities", util.LogValueFunc(func() any {
			return MapSummaryActivitiesToAbridged(activities)
		}))

	activityIds := make([]int64, len(activities))
	for i, activity := range activities {
		activityIds[i] = *activity.Id
	}
	streams, err := RetrieveStreams(stravaClient, activityIds)
	if err != nil {
		return nil, err
	}

	slog.Info("retrieved streams",
		"length", len(streams),
		"streams", util.LogValueFunc(func() any {
			vals := make([]map[string]any, len(activities))
			for i, s := range streams {
				vals[i] = map[string]any{}
				m, _ := s.ToMap()
				for k, v := range m {
					vMap := structs.Map(v)
					delete(vMap, "Data")
					vals[i][k] = vMap
				}
			}
			return vals
		}))

	zipped := make([]ActivityAndStream, len(activities))
	for i, a := range activities {
		a := a
		streamSet := streams[i]
		zipped[i] = ActivityAndStream{Activity: &a, StreamSet: streamSet}
	}

	return zipped, nil
}
