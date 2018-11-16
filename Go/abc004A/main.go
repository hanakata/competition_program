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

func solved(n int) int {
	solve := n * 2
	return solve
}

func main() {
	a := nextString()
	n, _ := strconv.Atoi(a)
	solve := solved(n)
	fmt.Println(solve)
	return
}
