package main

import (
	"testing"
)

type solvedTest struct {
	in1, in2, in3, out int
}

var solvedTests = []solvedTest{
	{12, 18, 11, 12},
}

func TestAbc019A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2, st.in3)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
