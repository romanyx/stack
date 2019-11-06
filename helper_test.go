package stack_test

import (
	"errors"

	"github.com/romanyx/stack"
)

var errOrigin = errors.New("origin")

func frameA() error {
	return stack.Errorf("frameA: %w", frameB())
}

func frameB() error {
	return stack.Errorf("frameB: %w", frameC())
}

func frameC() error {
	return errOrigin
}
