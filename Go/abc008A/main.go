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

func solved(s int, t int) int {
	solve := t - s + 1
	return solve
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextString()
	b := nextString()
	s, _ := strconv.Atoi(a)
	t, _ := strconv.Atoi(b)
	solve := solved(s, t)
	fmt.Println(solve)
	return
}
