// Package pipe provides utilities to pipe functions.
package pipe

type result[T any] struct {
	v   T
	err error
}

type evaluator[T any] interface {
	eval() (*result[T], func())
	registerFn(f func(T))
}

func evalResult[T any](v T, fn func(), fns []func(T)) (*result[T], func()) {
	cleanup := func() {
		for i := len(fns) - 1; i >= 0; i-- {
			fns[i](v)
		}
		fn()
	}
	return &result[T]{v, nil}, cleanup
}

type scalar[T any] struct {
	v        T
	cleanups []func(T)
}

func (s *scalar[T]) eval() (*result[T], func()) {
	return evalResult(s.v, func() {}, s.cleanups)
}

func (s *scalar[T]) registerFn(f func(T)) {
	s.cleanups = append(s.cleanups, f)
}

type selection[In, Out any] struct {
	parent   evaluator[In]
	fn       func(in In) (Out, error)
	cleanups []func(Out)
}

func (s *selection[In, Out]) eval() (*result[Out], func()) {
	r1, cleanup := s.parent.eval()
	if r1.err != nil {
		var zero Out
		return &result[Out]{zero, r1.err}, cleanup
	}
	r2, err := s.fn(r1.v)
	if err != nil {
		var zero Out
		return &result[Out]{zero, err}, cleanup
	}
	return evalResult(r2, cleanup, s.cleanups)
}

func (s *selection[In, Out]) registerFn(f func(Out)) {
	s.cleanups = append(s.cleanups, f)
}

func withError[In, Out any](f func(In) Out) func(In) (Out, error) {
	return func(v In) (Out, error) {
		return f(v), nil
	}
}

// Pipe represents the term of the pipeline.
type Pipe[T any] struct {
	next evaluator[T]
}

// Value returns a term.
func Value[T any](v T) *Pipe[T] {
	return &Pipe[T]{&scalar[T]{v, nil}}
}

func (p *Pipe[T]) To[T2 any](f func(T) T2) *Pipe[T2] {
	return &Pipe[T2]{
		next: &selection[T, T2]{p.next, withError(f), nil},
	}
}

func (p *Pipe[T]) Try[T2 any](f func(T) (T2, error)) *Pipe[T2] {
	return &Pipe[T2]{
		next: &selection[T, T2]{p.next, f, nil},
	}
}

// Defer registers f to cleanup T.
func (p *Pipe[T]) Defer(f func(v T)) *Pipe[T] {
	p.next.registerFn(f)
	return p
}

// Eval returns the result of the pipeline.
// If the pipeline gets an error, it stops the rest of the evaluations and returns that error along with the zero value of T.
func (p *Pipe[T]) Eval() (T, error) {
	r, cleanup := p.next.eval()
	defer cleanup()
	if r.err != nil {
		var zero T
		return zero, r.err
	}
	return r.v, nil
}
