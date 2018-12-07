package main

import (
	"testing"
)

type solvedTest struct {
	in1, in2, in3 int
	out           bool
}

var solvedTests = []solvedTest{
	{50, 60, 70, false},
	{60, 60, 70, true},
	{70, 60, 70, true},
	{80, 60, 70, false},
}

func TestAbc022A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2, st.in3)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
