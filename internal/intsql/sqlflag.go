package intsql

import (
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type SqlFlags interface {
	util.Flags
	NewStravaDatabase() (StravaDatabase, error)
}
