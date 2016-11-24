package failures

import "errors"

// Global errors
var (
	ErrWrongStartCmdCli = errors.New("Wrong command to use the app, use run or migrate")
)
