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
	a := n % 2
	if a == 0 {
		solve = n / 2
	} else {
		solve = n/2 + 1
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
