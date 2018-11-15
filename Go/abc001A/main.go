package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func solved(h1 int, h2 int) int {
	solve := h1 - h2
	return solve
}

func main() {
	a, b := nextLine(), nextLine()
	h1, _ := strconv.Atoi(a)
	h2, _ := strconv.Atoi(b)
	solve := solved(h1, h2)
	fmt.Println(solve)
	return
}
