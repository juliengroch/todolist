package failures

import "errors"

// Global errors
var (
	ErrBadFlagCli = errors.New("Bad argument for start flag")
)
