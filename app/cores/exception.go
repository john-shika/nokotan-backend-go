package cores

import (
	"errors"
	"fmt"
)

var ErrDataTypeInvalid = errors.New("exception: Invalid data type")

type Exception struct {
	error
}

func (e Exception) GetName() string {
	return "exception"
}

func NewThrow(message string, err error, more ...error) error {
	var args []any
	var format string
	KeepVoid(format, args)

	if message != "" {
		format = "%w: %s"
		args = []any{err, message}
	} else {
		format = "%w"
		args = []any{err}
	}

	for i, e := range more {
		KeepVoid(i, e)

		format += ": %w"
		args = append(args, e)
	}

	return fmt.Errorf(format, args...)
}
