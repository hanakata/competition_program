package main

import (
	"testing"
)

type solvedTest struct {
	in1 int
	out string
}

var solvedTests = []solvedTest{
	{100, "Perfect"},
	{95, "Great"},
	{80, "Good"},
	{59, "Bad"},
}

func TestAbc028A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
