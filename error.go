package stack

import (
	"errors"
	"runtime"
)

func (e *withStack) Error() string {
	return e.origin.Error()
}

func (e *withStack) Unwrap() error {
	return errors.Unwrap(e.origin)
}

func (e *withStack) StackTrace() []runtime.Frame {
	return e.stack.StackTrace()
}

type withStack struct {
	origin error
	*stack
}

// frame represents a program counter inside a stack frame.
// For historical reasons if Frame is interpreted as a uintptr
// its value represents the program counter + 1.
type frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f frame) pc() uintptr { return uintptr(f) - 1 }

// file returns the full path to the file that contains the
// function for this Frame's pc.
func (f frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

// line returns the line number of source code of the
// function for this Frame's pc.
func (f frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// name returns the name of this function, if known.
func (f frame) name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

func (f frame) runtime() runtime.Frame {
	rf := runtime.Frame{
		PC:   f.pc(),
		File: f.file(),
		Line: f.line(),
	}

	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return rf
	}

	rf.Func = fn
	rf.Function = fn.Name()
	rf.Entry = fn.Entry()

	return rf
}

// stack represents a stack of program counters.
type stack []uintptr

func (s *stack) StackTrace() []runtime.Frame {
	f := make([]runtime.Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = frame((*s)[i]).runtime()
	}
	return f
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
