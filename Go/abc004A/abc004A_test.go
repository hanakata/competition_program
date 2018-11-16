package main

import "testing"

type solvedTest struct {
	in, out int
}

var solvedTests = []solvedTest{
	{1000, 2000},
	{1000000, 2000000},
	{0, 0},
}

func TestAbc004A(t *testing.T) {

	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
