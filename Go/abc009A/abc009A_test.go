package main

import "testing"

type solvedTest struct {
	in  int
	out int
}

var solvedTests = []solvedTest{
	{2, 1},
	{5, 3},
	{1, 1},
}

func TestAbc009A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
