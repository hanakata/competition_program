package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func solved(a int) []int {
	s := []int{}
	if a >= 8 {
		s = append(s, 8)
		a -= 8
	}
	if a >= 4 {
		s = append(s, 4)
		a -= 4
	}
	if a >= 2 {
		s = append(s, 2)
		a -= 2
	}
	if a >= 1 {
		s = append(s, 1)
	}
	return s
}

func main() {
	a := nextString()
	n, _ := strconv.Atoi(a)
	solve := solved(n)
	fmt.Println(len(solve))
	for _, s := range solve {
		fmt.Println(s)
	}
	return
}
