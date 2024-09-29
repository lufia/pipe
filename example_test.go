package pipe_test

import (
	"fmt"
	"slices"
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

func ExampleValue() {
	p1 := pipe.Value("hello world").
		TryChain(require).
		Chain(tee).
		Chain(strings.ToUpper)
	p2 := pipe.From(p1, strings.Fields)
	p3 := pipe.From(p2, slices.Values)
	p4 := pipe.Each(p3, func(s string) string {
		return s + "!"
	})
	p5 := pipe.From(p4, slices.Collect)
	a, _ := p5.Eval()
	fmt.Println(a)
	// Output:
	// hello world
	// [HELLO! WORLD!]
}
