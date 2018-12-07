package main

import (
	"testing"
)

type solvedTest struct {
	in1, in2 string
	out      int
}

var solvedTests = []solvedTest{
	{"2", "3", 5},
	{"7", "0", 7},
	{"5", "3", 8},
}

func TestAbc023A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
