package main

import "testing"

type solvedTest struct {
	in1, in2, out int
}

var solvedTests = []solvedTest{
	{10, 5, 10},
	{15, 5, 15},
}

func TestAbc002A(t *testing.T) {

	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
