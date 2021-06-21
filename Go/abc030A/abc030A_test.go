package main

import (
	"testing"
)

type solvedTest struct {
	in1 float64
	in2 float64
	in3 float64
	in4 float64
	out string
}

var solvedTests = []solvedTest{
	{5, 2, 6, 3, "AOKI"},
	{100, 80, 100, 73, "TAKAHASHI"},
	{66, 30, 55, 25, "DRAW"},
}

func TestAbc030A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2, st.in3, st.in4)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
