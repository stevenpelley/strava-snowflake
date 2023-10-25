package intsnowflake

import (
	"flag"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
)

type SFFlags struct {
	configFileName string
	etlTableName   string
}

// prove it is Flags
var _ intsql.SqlFlags = &SFFlags{}

// NewStravaDatabase implements intsql.Flags.
func (f *SFFlags) NewStravaDatabase() (intsql.StravaDatabase, error) {
	d := New(f.configFileName, f.etlTableName)
	return &d, nil
}

func (f *SFFlags) InitFlags(fs *flag.FlagSet) error {
	return nil
}

func (f *SFFlags) PostProcessFlags(fs *flag.FlagSet) error {
	return nil
}
