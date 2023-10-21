package intsnowflake

import (
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type SFFlags struct {
}

// prove it is Flags
var _ util.Flags = &SFFlags{}

func (df *SFFlags) InitFlags() error {
	return nil
}

func (df *SFFlags) PostProcessFlags() error {
	return nil
}
