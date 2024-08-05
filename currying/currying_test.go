package currying

import (
	"strings"
	"testing"
)

func TestLast2(t *testing.T) {
	contains := Last2R1(strings.Contains)("seafood")
	if b := contains("foo"); b != true {
		t.Errorf(`contains("foo") = %t`, b)
	}
	if b := contains("bar"); b != false {
		t.Errorf(`contains("bar") = %t`, b)
	}
}

func TestFirst2(t *testing.T) {
	s := "seafood"
	containsFoo := First2R1(strings.Contains)("foo")
	if b := containsFoo(s); b != true {
		t.Errorf("containsFoo(%q) = %t", s, b)
	}
	containsBar := First2R1(strings.Contains)("bar")
	if b := containsBar(s); b != false {
		t.Errorf("containsBar(%q) = %t", s, b)
	}
}
