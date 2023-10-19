package intduckdb

import (
	"flag"

	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type DuckdbFlags struct {
	DbFileName string
}

// prove it is Flags
var _ util.Flags = &DuckdbFlags{}

func (df *DuckdbFlags) InitFlags() error {
	flag.StringVar(&df.DbFileName, "duckdbfile", "", "duckdb database file (empty for memory database)")
	return nil
}

func (df *DuckdbFlags) PostProcessFlags() error {
	return nil
}
