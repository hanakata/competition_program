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

func solved(a, b int) int {
	return a * b / 10
}

func main() {
	var i, solve int
	sc.Split(bufio.ScanWords)
	for i < 3 {
		a := nextString()
		b := nextString()
		n, _ := strconv.Atoi(a)
		m, _ := strconv.Atoi(b)
		solve += solved(n, m)
		i++
	}
	fmt.Println(solve)
	return
}
