package main

import (
	"testing"
)

type solvedTest struct {
	in1 string
	out string
}

var solvedTests = []solvedTest{
	{"0000", "0000"},
	{"1011", "0101"},
	{"1111", "0111"},
}

func TestMain(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1)
		if actual != st.out {
			t.Errorf("%v != %v", st.out, actual)
		}
	}
}
