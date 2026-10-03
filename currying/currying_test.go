package currying

import (
	"strconv"
	"strings"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestMake1to2(t *testing.T) {
	f := Make1to2(strconv.Atoi)
	n, err := f("12")
	gt.NoError(t, err).Required()
	gt.Number(t, n).Equal(12)

	n, err = f.Apply1("12")()
	gt.NoError(t, err).Required()
	gt.Number(t, n).Equal(12)

	n, err = f.Reverse()("12")
	gt.NoError(t, err).Required()
	gt.Number(t, n).Equal(12)
}

func TestMake2to1(t *testing.T) {
	f := Make2to1(strings.Contains)
	gt.Bool(t, f("seafood")("foo")).True()
	gt.Bool(t, f.Apply1("seafood")("foo")).True()
	gt.Bool(t, f.Apply2("seafood", "foo")()).True()
	gt.Bool(t, f.Reverse()("foo")("seafood")).True()
}

func TestMake3(t *testing.T) {
	f := Make3(func(v1, v2, v3 int) {
		gt.Number(t, v1).Equal(1)
		gt.Number(t, v2).Equal(2)
		gt.Number(t, v3).Equal(3)
	})
	f(1)(2)(3)
	f.Apply1(1)(2)(3)
	f.Apply2(1, 2)(3)
	f.Apply3(1, 2, 3)()
	f.Reverse()(3)(2)(1)
}
