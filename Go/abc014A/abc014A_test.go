package main

import "testing"

type solvedTest struct {
	in1, in2, out int
}

var solvedTests = []solvedTest{
	{7, 3, 2},
	{5, 5, 0},
	{1, 100, 99},
}

func TestAbc014A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
