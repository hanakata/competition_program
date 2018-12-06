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

func solved(n, m int) int {
	var s int
	a := n % m
	if a == 0 {
		s = a
	} else {
		s = m - a
	}
	return s
}

func main() {
	a := nextString()
	b := nextString()
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	solve := solved(n, m)
	fmt.Println(solve)
	return
}
