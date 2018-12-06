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

func solved(a string) string {
	s := "chokudai"
	if a == "1" {
		s = "ABC"
	}
	return s
}

func main() {
	a := nextString()
	solve := solved(a)
	fmt.Println(solve)
	return
}
