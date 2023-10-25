package util

import "flag"

type Flags interface {
	InitFlags(*flag.FlagSet) error
	PostProcessFlags(*flag.FlagSet) error
}
