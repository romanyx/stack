package stack

import (
	"errors"
	"runtime"
)

const defaultStackCap = 20

type withStack struct {
	origin error
	stack  []runtime.Frame
}

func (e *withStack) Error() string {
	return e.origin.Error()
}

func (e *withStack) Unwrap() error {
	return errors.Unwrap(e.origin)
}

func (e *withStack) StackTrace() []runtime.Frame {
	return e.stack
}

func trace(err error, offset int) *withStack {
	e := withStack{
		origin: err,
		stack:  make([]runtime.Frame, 0, defaultStackCap),
	}

	for {
		if f, ok := getFrame(offset); ok {
			e.stack = append(e.stack, f)
			offset++
			continue
		}

		break
	}

	return &e
}

func getFrame(caller int) (runtime.Frame, bool) {
	pc, file, line, ok := runtime.Caller(caller)
	if !ok {
		return runtime.Frame{}, false
	}

	f := runtime.Frame{
		PC:   pc,
		File: file,
		Line: line,
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return f, true
	}

	f.Func = fn
	f.Function = fn.Name()
	f.Entry = fn.Entry()

	return f, true
}
