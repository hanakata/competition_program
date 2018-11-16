package main

import "testing"

type solvedTest struct {
	in, out int
}

var solvedTests = []solvedTest{
	{6, 35000},
	{91, 460000},
}

func TestAbc003A(t *testing.T) {

	for _, st := range solvedTests {
		actual := solved(st.in)
		if actual != st.out {
			t.Errorf("%d != %d", st.out, actual)
		}
	}
}
