package errors

import (
	"fmt"
	"strings"
)

// Error - define Error as constant
type Error string

func (e Error) Error() string {
	return string(e)
}

// Is implements golang.org/pkg/errors/Is allowing a Error to check if it is the same even when wrapped
// This implementation only check the top most error
func (e Error) Is(target error) bool {
	return e.Error() == target.Error() || strings.HasPrefix(target.Error(), e.Error()+": ")
}

type wrappedError struct {
	cause error
	msg   string
}

func (w wrappedError) Error() string {
	if w.cause != nil {
		return fmt.Sprintf("%s: %v", w.msg, w.cause)
	}

	return w.msg
}

func (e Error) Wrap(err error) error {
	return wrappedError{cause: err, msg: string(e)}
}
