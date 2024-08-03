package pipe_test

import (
	"fmt"
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
	p := pipe.Value("hello world").
		TryChain(require).
		Chain(tee).
		Chain(strings.ToUpper)
	a, _ := pipe.From(p, strings.Fields).Eval()
	fmt.Println(a)
	// Output:
	// hello world
	// [HELLO WORLD]
}
