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
		Try(require).
		To(tee).
		To(strings.ToUpper).
		To(strings.Fields).
		To(slices.Values)
	p2 := pipe.Each(p1, func(s string) string { return s + "!" }).
		To(slices.Collect)
	a, _ := p2.Eval()
	fmt.Println(a)
	// Output:
	// hello world
	// [HELLO! WORLD!]
}
