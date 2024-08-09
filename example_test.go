package pipe_test

import (
	"fmt"
	"iter"
	"strings"

	"github.com/lufia/pipe"
)

func tee[T any](v T) T {
	fmt.Println(v)
	return v
}

func require[T ~string](v T) (T, error) {
	if len(v) == 0 {
		return "", fmt.Errorf("zero length")
	}
	return v, nil
}

// After Go 1.23 is released, slices.Values will replace this.
func values[T any](a []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, s := range a {
			if !yield(s) {
				break
			}
		}
	}
}

// After Go 1.23 is released, slices.Collect will replace this.
func collect[T any](seq iter.Seq[T]) []T {
	var a []T
	for v := range seq {
		a = append(a, v)
	}
	return a
}

func ExampleValue() {
	p1 := pipe.Value("hello world").
		TryChain(require).
		Chain(tee).
		Chain(strings.ToUpper)
	p2 := pipe.From(p1, strings.Fields)
	p3 := pipe.From(p2, values)
	p4 := pipe.Each(p3, func(s string) string {
		return s + "!"
	})
	p5 := pipe.From(p4, collect)
	a, _ := p5.Eval()
	fmt.Println(a)
	// Output:
	// hello world
	// [HELLO! WORLD!]
}
