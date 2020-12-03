package stack

import (
	"errors"
	"fmt"
	"runtime"
)

// Errorf saves stack trace and pass arguments to
// fmt.Errorf.
func Errorf(format string, a ...interface{}) error {
	return &withStack{
		fmt.Errorf(format, a...),
		callers(),
	}
}

// Origin returns unwrapped origin of the error.
func Origin(err error) error {
	if err == nil {
		return nil
	}

	for {
		e := err
		err = errors.Unwrap(err)
		if err == nil {
			return e
		}
	}
}

// Trace returns stack trace for error.
func Trace(err error) []runtime.Frame {
	var stack []runtime.Frame

	for {
		if err == nil {
			break
		}

		e, ok := err.(*withStack)
		if !ok {
			err = errors.Unwrap(err)
			continue
		}

		stack = e.StackTrace()
		err = errors.Unwrap(err)
	}

	if len(stack) == 0 {
		return nil
	}

	return stack
}
