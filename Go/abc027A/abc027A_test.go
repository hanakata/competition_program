package main

import (
	"testing"
)

type solvedTest struct {
	in1 string
	out string
}

var solvedTests = []solvedTest{
	{"1 1 2", "2"},
	{"1 2 1", "2"},
	{"2 1 1", "2"},
}

func TestAbc027A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
