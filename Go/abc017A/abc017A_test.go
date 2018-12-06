package main

import "testing"

type solvedTest struct {
	in1, in2, out int
}

var solvedTests = []solvedTest{
	{50, 7, 35},
	{40, 8, 32},
	{30, 9, 27},
}

func TestAbc017A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
