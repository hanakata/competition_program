package main

import (
	"testing"
)

type solvedTest struct {
	in1 string
	in2 int
	out string
}

var solvedTests = []solvedTest{
	{"abcdef", 1, "aa"},
	{"abcdef", 2, "ab"},
}

func TestAbc025A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
