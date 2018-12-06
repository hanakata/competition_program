package main

import "testing"

type solvedTest struct {
	in  string
	out int
}

var solvedTests = []solvedTest{
	{"A", 1},
	{"B", 2},
	{"C", 3},
	{"D", 4},
	{"E", 5},
}

func TestAbc013A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
