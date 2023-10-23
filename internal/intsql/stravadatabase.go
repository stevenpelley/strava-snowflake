package intsql

import (
	"database/sql"

	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type StravaDatabase interface {
	OpenDB() error
	DB() *sql.DB
	InitAndValidateSchema() error
	FilterKnownActivityIds(activityIds []int64) ([]int64, error)
	UploadActivityJson(activities []util.Jsonable) error
	MergeActivities() error
	Close() error
}
