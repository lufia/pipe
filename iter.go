package pipe

import (
	"iter"
)

type sequence[In, Out any] struct {
	parent   evaluator[iter.Seq[In]]
	fn       func(In) Out
	cleanups []func(iter.Seq[Out])
}

func (s *sequence[In, Out]) eval() (*result[iter.Seq[Out]], func()) {
	r, cleanup := s.parent.eval()
	if r.err != nil {
		return &result[iter.Seq[Out]]{nil, r.err}, cleanup
	}
	var seq iter.Seq[Out]
	seq = func(yield func(Out) bool) {
		for v := range r.v {
			if !yield(s.fn(v)) {
				break
			}
		}
	}
	return evalResult(seq, cleanup, s.cleanups)
}

func (s *sequence[In, Out]) registerFn(f func(iter.Seq[Out])) {
	s.cleanups = append(s.cleanups, f)
}

func Each[In, Out any](p *Pipe[iter.Seq[In]], f func(In) Out) *Pipe[iter.Seq[Out]] {
	return &Pipe[iter.Seq[Out]]{
		next: &sequence[In, Out]{p.next, f, nil},
	}
}
