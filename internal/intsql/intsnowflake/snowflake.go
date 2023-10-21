package intsnowflake

import (
	"database/sql"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type SFStrava struct {
	db *sql.DB
}

// prove it is a StravaDatabase
var _ intsql.StravaDatabase = &SFStrava{}

// OpenDB implements intsql.StravaDatabase.
func (sdb *SFStrava) OpenDB(dbFileName string) error {
	//config := gosnowflake.Config{}
	panic("unimplemented")
}

// GetExistingActivityIds implements intsql.StravaDatabase.
func (sdb *SFStrava) GetExistingActivityIds() (strava.IntSet, error) {
	panic("unimplemented")
}

// InitAndValidateSchema implements intsql.StravaDatabase.
func (sdb *SFStrava) InitAndValidateSchema() error {
	panic("unimplemented")
}

// UploadActivityJson implements intsql.StravaDatabase.
func (sdb *SFStrava) UploadActivityJson(activities []util.Jsonable) error {
	panic("unimplemented")
}

// MergeActivities implements intsql.StravaDatabase.
func (sdb *SFStrava) MergeActivities() error {
	// no merge necessary in snowflake
	return nil
}

// Close implements intsql.StravaDatabase.
func (sdb *SFStrava) Close() error {
	return sdb.db.Close()
}
