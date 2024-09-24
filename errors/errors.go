package dirioerr

import (
	"errors"
	"fmt"
)

var (
	ErrDirInvalidArguments = errors.New("invalid argumnetd to command line")

	ErrDirCannotRead = errors.New("cannot read the file, check file path and name")
)

func PrintError(directory string, err error) {
	fmt.Printf("%s : %s", directory, err.Error())
}
