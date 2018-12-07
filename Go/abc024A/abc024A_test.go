package main

import (
	"testing"
)

type solvedTest struct {
	in1, in2, in3, in4, in5, in6, out int
}

var solvedTests = []solvedTest{
	{100, 200, 50, 20, 40, 10, 3500},
	{400, 1000, 400, 21, 10, 10, 14000},
}

func TestAbc024A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2, st.in3, st.in4, st.in5, st.in6)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
