package main

import (
	"testing"
)

type solvedTest struct {
	in1 int
	out string
}

var solvedTests = []solvedTest{
	{10, "25"},
	{11, "30"},
}

func TestAbc026A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
