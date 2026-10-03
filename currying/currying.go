// Package currying provides utilities for currying.
package currying

//go:generate go run template.go

type Func0 func()

func (f Func0) Reverse() Func0 {
	return f
}

func Make0(f func()) Func0 {
	return f
}
