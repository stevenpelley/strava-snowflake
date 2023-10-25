package intsnowflake

import (
	"errors"
	"flag"
	"log/slog"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
)

type SFFlags struct {
	configFileName string
	// absolutely no attempt is made to sanitize this input.  Use caution.
	etlTableName string
}

// prove it is Flags
var _ intsql.SqlFlags = &SFFlags{}

// NewStravaDatabase implements intsql.Flags.
func (f *SFFlags) NewStravaDatabase() (intsql.StravaDatabase, error) {
	d := New(f.configFileName, f.etlTableName)
	return &d, nil
}

func (f *SFFlags) InitFlags(fs *flag.FlagSet) error {
	fs.StringVar(
		&f.configFileName,
		"configfilename",
		"",
		"json config file for snowflake.  Must match config.go")
	fs.StringVar(
		&f.etlTableName,
		"etltablename",
		"",
		"table into which to insert.  Must have a variant column named data.")
	return nil
}

func (f *SFFlags) PostProcessFlags(fs *flag.FlagSet) error {
	if f.configFileName == "" {
		return errors.New(`"-configfilename" required`)
	}
	if f.etlTableName == "" {
		return errors.New(`"-etlTableName" required`)
	}
	slog.Info("initialized snowflake inputs", "configfile", f.configFileName, "etltablename", f.etlTableName)
	return nil
}
