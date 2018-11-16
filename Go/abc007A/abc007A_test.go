package main

import "testing"

type solvedTest struct {
	in  int
	out int
}

var solvedTests = []solvedTest{
	{4, 3},
	{100, 99},
	{1, 0},
}

func TestAbc007A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
