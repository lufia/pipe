package pipe

import (
	"strconv"
	"strings"
	"testing"

	"golang.org/x/exp/constraints"
)

func wantEval[T comparable](t *testing.T, x T, p *Pipe[T]) {
	t.Helper()
	v, err := p.Eval()
	if err != nil {
		t.Errorf("Eval(): %v", err)
		return
	}
	if v != x {
		t.Errorf("Eval() = %v; want %v", v, x)
	}
}

func TestValue(t *testing.T) {
	wantEval(t, 10, Value(10))
}

func TestPipeChain(t *testing.T) {
	add2 := func(n int) int {
		return n + 2
	}
	wantEval(t, 12, Value(10).Chain(add2))
}

func TestPipeTryChain(t *testing.T) {
	add10 := func(n int) (int, error) {
		return n + 10, nil
	}
	wantEval(t, 20, Value(10).TryChain(add10))
}

func TestFrom(t *testing.T) {
	wantEval(t, "10", From(Value(10), strconv.Itoa))
}

func TestPipeTryFrom(t *testing.T) {
	wantEval(t, 10, TryFrom(Value("10"), strconv.Atoi))
}

func TestPipe(t *testing.T) {
	plus3 := func(n int) int { return n + 3 }
	times10 := func(n int) int { return n * 10 }
	v := Value(10)
	v.Pipe(plus3)(times10)
	wantEval(t, 130, v)
}

func writerSeq[In any, Out constraints.Integer](f func(In) (Out, error)) func(In) func(Out) (Out, error) {
	return func(v In) func(Out) (Out, error) {
		return func(n Out) (Out, error) {
			r, err := f(v)
			return n + r, err
		}
	}
}

func TestPipeErr(t *testing.T) {
	var w strings.Builder
	write := writerSeq(w.WriteString)
	p := Value(0).TryChain(write("hello")).TryChain(write("world"))
	wantEval(t, 10, p)
	if s := w.String(); s != "helloworld" {
		t.Errorf("Value = %s; want helloworld", s)
	}
}
