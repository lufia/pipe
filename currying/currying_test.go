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

func TestAll3(t *testing.T) {
	replace := All3R1(strings.ReplaceAll)
	s := replace("oink oink oink")("oink")("moo")
	if s != "moo moo moo" {
		t.Errorf("replace() = %s", s)
	}
}
