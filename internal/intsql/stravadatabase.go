package intsql

import (
	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type StravaDatabase interface {
	OpenDB(dbFileName string) error
	InitAndValidateSchema() error
	GetExistingActivityIds() (strava.IntSet, error)
	UploadActivityJson(activities []util.Jsonable) error
	MergeActivities() error
	Close() error
}
