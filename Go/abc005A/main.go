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
	solve := y / x
	return solve
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextString()
	b := nextString()
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	solve := solved(x, y)
	fmt.Println(solve)
	return
}
