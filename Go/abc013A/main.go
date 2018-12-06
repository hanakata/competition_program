package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func solved(a string) int {
	s := [...]string{"A", "B", "C", "D", "E"}
	var i int
	for _, n := range s {
		if n == a {
			return i + 1
		}
		i++
	}
	return i
}

func main() {
	a := nextString()
	solve := solved(a)
	fmt.Println(solve)
	return
}
