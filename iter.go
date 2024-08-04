//go:build ignore

package pipe

import (
	"iter"
)

type sequence[T any] struct {
}

type PipeSeq[SIn ~iter.Seq[EIn], EIn, EOut any] struct {
	next evaluator[SIn]
	fn   func(EIn) EOut
}

func Seq[SIn ~iter.Seq[EIn], EIn, EOut any](p *Pipe[S], f func(EIn) EOut) *PipeSeq[iter.Seq[EOut]] {
	return &PipeSeq[SIn, iter.Seq[E]]{
		next: p.next,
		fn:   f,
	}
}

func (p *PipeSeq[SIn, EIn, EOut]) Chain(f func(EIn) EOut) *PipeSeq[EOut] {
	seq := func(yield func(EOut) bool) {
	}
	return &PipeSeq[SIn, EIn, EOut]{
		next: seq,
		fn:   f,
	}
}

func (p *PipeSeq[SIn, EIn, EOut]) Eval() *PipeSeq[SIn, EIn, EOut] {
	seq := func(yield func(EOut) bool) {
		seq, cleanup := p.next.eval()
		defer cleanup()
		for v := range seq {
			yield(p.fn(v))
		}
	}
	return &PipeSeq[iter.Seq[EOut]]{}
}
