package stack

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkErrorf(b *testing.B) {
	for _, frames := range []int{5, 10, 20, 40} {
		suffix := fmt.Sprintf("%d-frames", frames)
		b.Run(suffix, func(b *testing.B) {
			err := Errorf("annotate: %w", errors.New("original"))
			// Reduce by number of parent frames in order to have a correct depth.
			depth := frames - len(Trace(err))
			if depth < 1 {
				panic("number of frames is negative")
			}

			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				addFrames(depth, errors.New("origin"))
			}
		})
	}
}

func addFrames(depth int, err error) error {
	if depth <= 1 {
		return Errorf("depth: %w", err)
	}
	return addFrames(depth-1, err)
}
