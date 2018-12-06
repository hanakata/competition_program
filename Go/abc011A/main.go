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
	var solve int
	if n == 12 {
		solve = 1
	} else {
		solve = n + 1
	}
	return solve
}

func main() {
	a := nextString()
	n, _ := strconv.Atoi(a)
	solve := solved(n)
	fmt.Println(solve)
	return
}
