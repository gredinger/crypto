package crypto

import "testing"

// https://golang.org/doc/code.html dude...thanks for that
func TestCeasarShiftLeft(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"ABC", "XYZ"},
		{"SALETHESHIPS", "PXIBQEBPEFMP"},
	}
	for _, c := range cases {
		got := CeasarShiftLeft(c.in)
		if got != c.want {
			t.Errorf("Ceasar Shift Left(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// Hehe,  SALE ;)

func TestCeasarShiftRight(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"XYZ", "ABC"},
		{"PXIBQEBPEFMP", "SALETHESHIPS"},
	}
	for _, c := range cases {
		got := CeasarShiftRight(c.in)
		if got != c.want {
			t.Errorf("Ceasar Shift Left(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
