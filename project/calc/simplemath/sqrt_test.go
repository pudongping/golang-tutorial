package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	v := Sqrt(9)
	if v != 3 {
		t.Errorf("Sqrt(9) failed. Got %v, expected 3.", v)
	}
}