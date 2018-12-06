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

func solved(a, b string) string {
	var s string
	if len(a) < len(b) {
		s = b
	} else {
		s = a
	}
	return s
}

func main() {
	a := nextString()
	b := nextString()
	solve := solved(a, b)
	fmt.Println(solve)
	return
}
