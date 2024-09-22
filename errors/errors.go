package dirioerr

import "errors"

var (
	ErrDirDonotHaveReadPermission = errors.New("directory read operatioin is not permitted")
	
	ErrDirInvalidArguments = errors.New("Invalid argumnetd to command line")
)
