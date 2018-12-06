package main

import "testing"

type solvedTest struct {
	in1, in2 string
	out      string
}

var solvedTests = []solvedTest{
	{"1", "2", "2 1"},
	{"3", "4", "4 3"},
	{"5", "5", "5 5"},
	{"13", "69", "69 13"},
}

func TestAbc012A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%s != %s", st.out, actual)
		}
	}
}
