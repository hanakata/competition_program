package main

import "testing"

type solvedTest struct {
	in1, in2 int
	out      string
}

var solvedTests = []solvedTest{
	{1, 1, "YES"},
	{2, 29, "NO"},
	{12, 6, "YES"},
}

func TestAbc016A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%s != %s", st.out, actual)
		}
	}
}
