// Package pipe provides utilities to pipe functions.
package pipe

type result[T any] struct {
	v   T
	err error
}

type evaluator[T any] interface {
	eval() *result[T]
}

type scalar[T any] struct {
	v T
}

func (s *scalar[T]) eval() *result[T] {
	return &result[T]{s.v, nil}
}

type selection[In, Out any] struct {
	parent evaluator[In]
	fn     func(v In) Out
}

func (s *selection[In, Out]) eval() *result[Out] {
	r := s.parent.eval()
	if r.err != nil {
		var zero Out
		return &result[Out]{zero, r.err}
	}
	return &result[Out]{s.fn(r.v), nil}
}

type selection2[In, Out any] struct {
	parent evaluator[In]
	fn     func(in In) (Out, error)
}

func (s *selection2[In, Out]) eval() *result[Out] {
	r1 := s.parent.eval()
	if r1.err != nil {
		var zero Out
		return &result[Out]{zero, r1.err}
	}
	r2, err := s.fn(r1.v)
	return &result[Out]{r2, err}
}

// Pipe represents the term of the pipeline.
type Pipe[T any] struct {
	next evaluator[T]
}

// Value returns a term.
func Value[T any](v T) *Pipe[T] {
	return &Pipe[T]{&scalar[T]{v}}
}

type Builder[T any] func(f func(T) T) Builder[T]

func (p *Pipe[T]) Pipe(f func(v T) T) Builder[T] {
	var add Builder[T]
	add = func(g func(T) T) Builder[T] {
		next := p.next
		p.next = &selection[T, T]{next, g}
		return add
	}
	add(f)
	return add
}

func (p *Pipe[T]) Chain(f func(v T) T) *Pipe[T] {
	return &Pipe[T]{
		next: &selection[T, T]{p.next, f},
	}
}

func (p *Pipe[T]) TryChain(f func(v T) (T, error)) *Pipe[T] {
	return &Pipe[T]{
		next: &selection2[T, T]{p.next, f},
	}
}

// Defer registers f to cleanup T.
func (p *Pipe[T]) Defer(f func(v T)) *Pipe[T] {
	return p
}

// Eval returns the result of the pipeline.
// If the pipeline gets an error, it stops the rest of the evaluations and returns that error along with the zero value of T.
func (p *Pipe[T]) Eval() (T, error) {
	r := p.next.eval()
	if r.err != nil {
		var zero T
		return zero, r.err
	}
	return r.v, nil
}

// From is like [Pipe.Chain] except f returns different type.
func From[In, Out any](p *Pipe[In], f func(In) Out) *Pipe[Out] {
	return &Pipe[Out]{
		next: &selection[In, Out]{p.next, f},
	}
}

// TryFrom is like [Pipe.TryChain] except f returns different type.
func TryFrom[In, Out any](p *Pipe[In], f func(In) (Out, error)) *Pipe[Out] {
	return &Pipe[Out]{
		next: &selection2[In, Out]{p.next, f},
	}
}

/*
If Go supports the type parameter on method, we will add pipe.To method.

  pipe.Value(auth()).To(tokenFrom).To(fetch)
*/
