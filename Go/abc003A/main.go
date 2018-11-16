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
	var m int
	for i := 0; i <= n; i++ {
		m += i * 10000
	}
	solve := m / n
	return solve
}

func main() {
	a := nextString()
	n, _ := strconv.Atoi(a)
	solve := solved(n)
	fmt.Println(solve)
	return
}
