package main

import "testing"

type solvedTest struct {
	in1, in2, out int
}

var solvedTests = []solvedTest{
	{4, 8, 2},
	{4, 7, 1},
	{4, 3, 0},
}

func TestAbc005A(t *testing.T) {

	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
