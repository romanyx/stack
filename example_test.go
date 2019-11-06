package stack_test

import (
	"errors"
	"fmt"

	"github.com/romanyx/stack"
)

func ExampleTrace() {
	err := frameA()
	for _, frame := range stack.Trace(err) {
		fmt.Printf("%s:%d %s()\n", frame.File, frame.Line, frame.Function)
	}

	// Example output:
	// github.com/romanyx/stack/helper_test.go:16 github.com/romanyx/stack_test.frameB()
	// github.com/romanyx/stack/helper_test.go:12 github.com/romanyx/stack_test.frameA()
	// github.com/romanyx/stack/example_test.go:10 github.com/romanyx/stack_test.ExampleTrance()
	// /usr/local/go/src/testing/run_example.go:62 testing.runExample()
	// /usr/local/go/src/testing/example.go:44 testing.runExamples()
	// /usr/local/go/src/testing/testing.go:1118 testing.(*M).Run()
	// _testmain.go:52 main.main()
	// /usr/local/go/src/runtime/proc.go:203 runtime.main()
	// /usr/local/go/src/runtime/asm_amd64.s:1357 runtime.goexit()
}

func ExampleOrigin() {
	err := frameA()
	fmt.Println(stack.Origin(err))

	// Output:
	// origin
}

func ExampleErrorf() {
	err := stack.Errorf("annotated: %w", errOrigin)
	fmt.Println(errors.Is(err, errOrigin))

	// Output:
	// true
}
