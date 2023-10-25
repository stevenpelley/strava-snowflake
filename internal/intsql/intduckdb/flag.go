package intduckdb

import (
	"flag"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
)

type DuckdbFlags struct {
	DbFileName string
}

// prove it is Flags
var _ intsql.SqlFlags = &DuckdbFlags{}

// NewStravaDatabase implements intsql.Flags.
func (f *DuckdbFlags) NewStravaDatabase() (intsql.StravaDatabase, error) {
	d := New(f.DbFileName)
	return &d, nil
}

func (df *DuckdbFlags) InitFlags(fs *flag.FlagSet) error {
	fs.StringVar(&df.DbFileName, "duckdbfile", "", "duckdb database file (empty for memory database)")
	return nil
}

func (df *DuckdbFlags) PostProcessFlags(fs *flag.FlagSet) error {
	return nil
}
