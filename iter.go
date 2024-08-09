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
		for in := range r.v {
			if !yield(s.fn(in)) {
				break
			}
		}
	}
	return evalResult(seq, cleanup, s.cleanups)
}

func (s *sequence[In, Out]) registerFn(f func(iter.Seq[Out])) {
	s.cleanups = append(s.cleanups, f)
}

type sequence2[KIn, VIn, KOut, VOut any] struct {
	parent   evaluator[iter.Seq2[KIn, VIn]]
	fn       func(KIn, VIn) (KOut, VOut)
	cleanups []func(iter.Seq2[KOut, VOut])
}

func (s *sequence2[KIn, VIn, KOut, VOut]) eval() (*result[iter.Seq2[KOut, VOut]], func()) {
	r, cleanup := s.parent.eval()
	if r.err != nil {
		return &result[iter.Seq2[KOut, VOut]]{nil, r.err}, cleanup
	}
	var seq iter.Seq2[KOut, VOut]
	seq = func(yield func(KOut, VOut) bool) {
		for k, v := range r.v {
			if !yield(s.fn(k, v)) {
				break
			}
		}
	}
	return evalResult(seq, cleanup, s.cleanups)
}

func (s *sequence2[KIn, VIn, KOut, VOut]) registerFn(f func(iter.Seq2[KOut, VOut])) {
	s.cleanups = append(s.cleanups, f)
}

func Each[In, Out any](p *Pipe[iter.Seq[In]], f func(In) Out) *Pipe[iter.Seq[Out]] {
	return &Pipe[iter.Seq[Out]]{
		next: &sequence[In, Out]{p.next, f, nil},
	}
}

func TryEach[In, Out any](p *Pipe[iter.Seq[In]], f func(In) (Out, error)) *Pipe[iter.Seq2[Out, error]] {
	return &Pipe[iter.Seq2[Out, error]]{
		next: &sequence2[In, In, Out, error]{}, // TODO: implement
	}
}

func Each2[KIn, VIn, KOut, VOut any](p *Pipe[iter.Seq2[KIn, VIn]], f func(KIn, VIn) (KOut, VOut)) *Pipe[iter.Seq2[KOut, VOut]] {
	return &Pipe[iter.Seq2[KOut, VOut]]{
		next: &sequence2[KIn, VIn, KOut, VOut]{p.next, f, nil},
	}
}
