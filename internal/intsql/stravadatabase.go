package intsql

import (
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type StravaDatabase interface {
	OpenDB(dbFileName string) error
	InitAndValidateSchema() error
	FilterKnownActivityIds(activityIds []int64) ([]int64, error)
	UploadActivityJson(activities []util.Jsonable) error
	MergeActivities() error
	Close() error
}
