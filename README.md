# Annotation with stack trace for go1.13

[![GoDoc](https://godoc.org/github.com/romanyx/stack?status.svg)](https://godoc.org/github.com/romanyx/stack)
[![Report](https://goreportcard.com/badge/github.com/romanyx/stack)](https://goreportcard.com/report/github.com/romanyx/stack)
[![Build Status](https://travis-ci.com/romanyx/stack.svg?branch=master)](https://travis-ci.com/romanyx/stack)

Go 1.13 contains support for [error wrapping](https://blog.golang.org/go1.13-errors). Now you can add additional information to an error by wrapping it using the new `%w` verb at `fmt.Errorf` and examine such errors using `errors.Is` and `errors.As`. If you also want to save a stack trace of an error instead of `fmt.Errorf` use `stack.Errorf` which is compatible with `errors.Is` and `errors.As` and also gives the ability to get a stack trace of the error using `stack.Trace` function which will return `[]runtime.Frame`.

* **Import**.

``` go
import "github.com/romanyx/stack
```

* **Annotate error**.

``` go
func example() error {
	if err := call(); err != nil {
		return stack.Errorf("call: %w", err)
	}

	return nil
}
```

* **Print original error**.

``` go
stack.Origin(err)
```

* **Iterate through stack trace**.

``` go
for _, frame := range stack.Trace(err) {
	fmt.Printf("%s:%d %s()\n", frame.File, frame.Line, frame.Function)
}
```

