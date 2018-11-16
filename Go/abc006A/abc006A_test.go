package main

import "testing"

type solvedTest struct {
	in  int
	out string
}

var solvedTests = []solvedTest{
	{2, "NO"},
	{9, "YES"},
	{3, "YES"},
}

func TestAbc006A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%s != %s", st.out, actual)
		}
	}
}
