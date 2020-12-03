package stack_test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/romanyx/stack"
)

func TestTrace(t *testing.T) {
	tt := []struct {
		name   string
		in     error
		expect []runtime.Frame
	}{
		{
			name:   "nil err",
			in:     nil,
			expect: nil,
		},
		{
			name:   "annotated",
			in:     fmt.Errorf("annotate: %v", errOrigin),
			expect: nil,
		},
		{
			name: "wrapped",
			in:   frameA(),
			expect: []runtime.Frame{
				{
					Function: "github.com/romanyx/stack_test.frameB",
				},
				{
					Function: "github.com/romanyx/stack_test.frameA",
				},
				{
					Function: "github.com/romanyx/stack_test.TestTrace",
				},
				{Function: "testing.tRunner"},
				{Function: "runtime.goexit"},
			},
		},
		{
			name: "fmt wrapped",
			in:   fmt.Errorf("fmt: %w", frameA()),
			expect: []runtime.Frame{
				{
					Function: "github.com/romanyx/stack_test.frameB",
				},
				{
					Function: "github.com/romanyx/stack_test.frameA",
				},
				{
					Function: "github.com/romanyx/stack_test.TestTrace",
				},
				{Function: "testing.tRunner"},
				{Function: "runtime.goexit"},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			stk := stack.Trace(tc.in)
			if len(stk) != len(tc.expect) {
				t.Errorf(
					"expected: %d frames got: %d",
					len(tc.expect),
					len(stk),
				)
				return
			}

			for i, frame := range stk {
				compareFrames(t, tc.expect[i], frame)
			}
		})
	}

}

func compareFrames(t *testing.T, expect, got runtime.Frame) {
	t.Helper()

	if expect.Function != got.Func.Name() {
		t.Errorf(
			"expected: %s function name got: %s",
			expect.Function,
			got.Func.Name(),
		)
	}
}

func TestOrigin(t *testing.T) {
	tt := []struct {
		name       string
		in, expect error
	}{
		{
			name:   "nil err",
			in:     nil,
			expect: nil,
		},
		{
			name:   "wrapped",
			in:     stack.Errorf("wrapped: %w", errOrigin),
			expect: errOrigin,
		},
		{
			name: "fmt wrapped",
			in: fmt.Errorf(
				"fmt: %w",
				stack.Errorf("wrapped: %w", errOrigin),
			),
			expect: errOrigin,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := stack.Origin(tc.in)
			if !errors.Is(got, tc.expect) {
				t.Errorf(
					"unexpected origin: %v expected: %v",
					got,
					tc.expect,
				)
			}
		})
	}
}
