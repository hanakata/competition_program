package main

import "testing"

type solvedTest struct {
	in  string
	out string
}

var solvedTests = []solvedTest{
	{"chokudai", "chokudaipp"},
	{"sanagi", "sanagipp"},
}

func TestAbc010A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%s != %s", st.out, actual)
		}
	}
}
