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

type Builder[T any] func(f func(T) T) Builder[T]

func (p *Pipe[T]) Pipe(f func(v T) T) Builder[T] {
	var add Builder[T]
	add = func(g func(T) T) Builder[T] {
		next := p.next
		p.next = &selection[T, T]{next, withError(g), nil}
		return add
	}
	add(f)
	return add
}

func (p *Pipe[T]) Chain(f func(v T) T) *Pipe[T] {
	return &Pipe[T]{
		next: &selection[T, T]{p.next, withError(f), nil},
	}
}

func (p *Pipe[T]) TryChain(f func(v T) (T, error)) *Pipe[T] {
	return &Pipe[T]{
		next: &selection[T, T]{p.next, f, nil},
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

// From is like [Pipe.Chain] except f returns different type.
func From[In, Out any](p *Pipe[In], f func(In) Out) *Pipe[Out] {
	return &Pipe[Out]{
		next: &selection[In, Out]{p.next, withError(f), nil},
	}
}

// TryFrom is like [Pipe.TryChain] except f returns different type.
func TryFrom[In, Out any](p *Pipe[In], f func(In) (Out, error)) *Pipe[Out] {
	return &Pipe[Out]{
		next: &selection[In, Out]{p.next, f, nil},
	}
}

/*
If Go supports the type parameter on method, we will add pipe.To method.

  pipe.Value(auth()).To(tokenFrom).To(fetch)
*/
