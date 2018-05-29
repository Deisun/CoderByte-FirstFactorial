package main

import (
	"testing"
)

type testpair struct {
	number int64
	factorial int64
}

func TestFirstFactorial(t *testing.T) {
	tests := []testpair{
		{0,1},
		{1,1},
		{2,2},
		{3,6},
		{4,24},
		{15,1307674368000},
	}

	for _,v  := range tests {
		f := FirstFactorial(v.number)

		if f != v.factorial {
			t.Errorf("Factorial was incorrect, got: %d, expected: %d", f, v.factorial)
		}
	}
}
