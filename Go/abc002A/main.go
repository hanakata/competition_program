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

func solved(x int, y int) int {
	var solve int
	if x < y {
		solve = y
	} else {
		solve = x
	}
	return solve
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextString()
	m := nextString()
	x, _ := strconv.Atoi(n)
	y, _ := strconv.Atoi(m)
	solve := solved(x, y)
	fmt.Println(solve)
	return
}
