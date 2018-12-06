package main

import "testing"

type solvedTest struct {
	in  int
	out int
}

var solvedTests = []solvedTest{
	{1, 2},
	{11, 12},
	{12, 1},
}

func TestAbc011A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
