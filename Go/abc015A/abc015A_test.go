package main

import "testing"

type solvedTest struct {
	in1, in2, out string
}

var solvedTests = []solvedTest{
	{"isuruu", "isleapyear", "isleapyear"},
	{"ttttiiiimmmmeeee", "time", "ttttiiiimmmmeeee"},
}

func TestAbc015A(t *testing.T) {
	for _, st := range solvedTests {
		actual := solved(st.in1, st.in2)
		if actual != st.out {
			t.Errorf("%s != %s", st.out, actual)
		}
	}
}
