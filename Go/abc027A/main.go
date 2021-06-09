package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func solved(a string) string {
	slice := strings.Split(a, " ")
	if slice[0] == slice[1] {
		return slice[2]
	}
	if slice[0] == slice[2] {
		return slice[1]
	}
	return slice[0]
}

func main() {
	a := nextString()
	solve := solved(a)
	fmt.Println(solve)
	return
}
