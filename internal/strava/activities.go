package strava

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/fatih/structs"
	"github.com/stevenpelley/strava-snowflake/internal/util"
	"github.com/stevenpelley/strava3golang"
	"golang.org/x/sync/errgroup"
)

type IntSet map[int64]struct{}

type ActivitiesConfig struct {
	StravaClient              *strava3golang.APIClient
	StartTime                 time.Time
	EndTime                   time.Time
	GetStreamsConcurrency     int
	ActivitiesTimeoutDuration time.Duration
	StreamsTimeoutDuration    time.Duration
	ActivityIdsToIgnore       IntSet
	ActivitiesPerPage         int
	PreStreamSleep            time.Duration
}

const activitiesPageSize int32 = 30

func RetrieveActivities(config *ActivitiesConfig) (
	[]strava3golang.SummaryActivity, error) {
	startUnixTime := config.StartTime.UTC().Unix()
	endUnixTime := config.EndTime.UTC().Unix()
	after := int32(startUnixTime)
	before := int32(endUnixTime)

	var currentPage int32 = 1
	var itemsInPage int = 0
	allActivities := make([]strava3golang.SummaryActivity, 0)
	ctx, cancelFunc := context.WithTimeout(
		context.Background(), config.ActivitiesTimeoutDuration)
	defer cancelFunc()

	for currentPage == 1 || int32(itemsInPage) == activitiesPageSize {
		slog.Debug("getting activity page", "current_page", currentPage)
		activities, _, err := config.StravaClient.ActivitiesAPI.GetLoggedInAthleteActivities(
			ctx).
			After(after).
			Before(before).
			PerPage(activitiesPageSize).
			Page(currentPage).
			Execute()
		if err != nil {
			return nil, fmt.Errorf("retrieving activity page: %w", err)
		}
		allActivities = append(allActivities, activities...)
		itemsInPage = len(activities)
		currentPage++
	}

	return allActivities, nil
}

// retrieve the corresponding StreamSet for each activity in activityIds
// note that when an error is returned the returned stream may be partially complete
func RetrieveStreams(config *ActivitiesConfig, activityIds []int64) ([]*strava3golang.StreamSet, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.StreamsTimeoutDuration)
	defer cancelFunc()
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(config.GetStreamsConcurrency)
	streams := make([]*strava3golang.StreamSet, len(activityIds))
	for i, activityId := range activityIds {
		i, activityId := i, activityId // bind to iteration-specific variable for closure
		g.Go(func() error {
			select {
			case val, ok := <-ctx.Done():
				if ok {
					log.Panicf("unexpected value from context Done channel: %v", val)
				} else {
					return fmt.Errorf("context is Done at beginning of task: %v", ctx.Err())
				}
			default:
				// not yet cancelled, proceed without blocking
			}

			if config.PreStreamSleep > 0 {
				slog.Debug("sleeping prior to retrieving stream",
					"activity_id", activityId, "activity_index", i)
				time.Sleep(config.PreStreamSleep)
			}
			slog.Debug("retrieving stream", "activity_id", activityId, "activity_index", i)
			streamSet, _, err := config.StravaClient.StreamsAPI.GetActivityStreams(ctx, activityId).
				KeyByType(true).
				Keys(AllStreamKindsString).
				Execute()
			if err != nil {
				slog.Warn("error retrieving stream", "activity_id", activityId, "activity_index", i, "error", err)
				return fmt.Errorf(
					"retrieving activity stream. id: %v. activity index: %v. error: %w",
					activityId, i, err)
			}
			slog.Debug("retrieved stream successfully", "activity_id", activityId, "activity_index", i)
			streams[i] = streamSet
			return nil
		})
	}
	err := g.Wait()
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

func MapActivitiesToAbridged(
	activities []Activity) []AbridgedSummaryActivity {
	vals := make([]AbridgedSummaryActivity, len(activities))
	for i, a := range activities {
		vals[i] = GetAbridgedActivity(a.summaryActivity)
	}
	return vals
}

type ActivityAndStream struct {
	Activity  *strava3golang.SummaryActivity
	StreamSet *strava3golang.StreamSet
}

func (aas *ActivityAndStream) ToJson() string {
	return util.MarshalOrPanic(aas)
}

type Activity struct {
	summaryActivity *strava3golang.SummaryActivity
}

func (a Activity) GetId() int64 {
	return *a.summaryActivity.Id
}

func getActivities(config *ActivitiesConfig) ([]Activity, error) {
	activities, err := RetrieveActivities(config)
	if err != nil {
		return nil, fmt.Errorf("retrieving activities: %w", err)
	}

	results := make([]Activity, len(activities))
	for i := 0; i < len(activities); i++ {
		results[i] = Activity{&activities[i]}
	}

	slog.Info("retrieved activities",
		"length", len(activities),
		"activities", util.LogValueFunc(func() any {
			return MapActivitiesToAbridged(results)
		}))

	return results, nil
}

func getStreams(config *ActivitiesConfig, activities []Activity) ([]util.Jsonable, error) {
	activityIds := make([]int64, len(activities))
	for i, activity := range activities {
		activityIds[i] = activity.GetId()
	}
	streams, err := RetrieveStreams(config, activityIds)
	if err != nil {
		slog.Warn(
			"error retrieving streams.  Attempting to create partial result",
			"error", err)
	}

	slog.Info("retrieved streams",
		"length", len(streams),
		"streams", util.LogValueFunc(func() any {
			vals := make([]map[string]any, len(activities))
			for i, s := range streams {
				vals[i] = map[string]any{}
				if s != nil {
					m, _ := s.ToMap()
					for k, v := range m {
						vMap := structs.Map(v)
						delete(vMap, "Data")
						vals[i][k] = vMap
					}
				}
			}
			return vals
		}))

	slog.Info("zipping result streams", "untruncated activities size", len(streams))
	zipped := make([]util.Jsonable, 0, len(activities))
	for i := 0; i < len(activities); i++ {
		activity := activities[i].summaryActivity
		stream := streams[i]
		if stream != nil {
			zipped = append(zipped, &ActivityAndStream{Activity: activity, StreamSet: stream})
		}
	}
	slog.Info("zipped result streams", "truncated activities size", len(zipped))

	// may have been a partial result so return previous error
	return zipped, err
}

// accepts a list of activity ids to consider.  Returns the list of activity
// ids whose streams should be downloaded, or an error
// it is expected that the returned ids are in the same relative order
// as the inputs
type ActivityFilter func([]int64) ([]int64, error)

func filterActivities(activities []Activity, filter ActivityFilter) ([]Activity, error) {
	activityIds := make([]int64, len(activities))
	for i, v := range activities {
		activityIds[i] = v.GetId()
	}
	activityIds, err := filter(activityIds)
	if err != nil {
		return nil, fmt.Errorf("FilterActivities: %w", err)
	}
	activityIdsSet := make(map[int64]struct{})
	for _, activityId := range activityIds {
		activityIdsSet[activityId] = struct{}{}
	}

	// expand ids back to activities
	results := make([]Activity, 0)
	for _, activity := range activities {
		_, ok := activityIdsSet[*activity.summaryActivity.Id]
		if ok {
			results = append(results, activity)
		}
	}

	slog.Info("filtered activities",
		"length", len(activities),
		"activities", util.LogValueFunc(func() any {
			return MapActivitiesToAbridged(results)
		}))

	return results, nil
}

// note that this may return a partial result even when error is not nil
func GetActivitiesAndStreams(config *ActivitiesConfig, filter ActivityFilter) (
	[]util.Jsonable, error) {
	activities, err := getActivities(config)
	if err != nil {
		return nil, err
	}

	activities, err = filterActivities(activities, filter)
	if err != nil {
		return nil, err
	}

	return getStreams(config, activities)
}
