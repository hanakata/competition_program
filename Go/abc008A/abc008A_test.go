package main

import "testing"

type solvedTest struct {
	in1 int
	in2 int
	out int
}

var solvedTests = []solvedTest{
	{4, 7, 4},
	{1, 1, 1},
}

func TestAbc008A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
