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

func solved(a, b, c int) [3]int {
	var i, j int
	var r = [3]int{a, b, c}
	var s = [3]int{1, 1, 1}
	for i < 3 {
		j = 0
		for j < 3 {
			if r[i] < r[j] {
				s[i] = s[i] + 1
			}
			j++
		}
		i++
	}
	return s
}

func main() {
	a := nextString()
	b := nextString()
	c := nextString()
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	l, _ := strconv.Atoi(c)
	solve := solved(n, m, l)
	for _, s := range solve {
		fmt.Println(s)
	}
	return
}
