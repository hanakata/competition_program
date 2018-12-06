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

func solved(a, b int) string {
	var i int
	i = a % b
	if i == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextString()
	b := nextString()
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	solve := solved(n, m)
	fmt.Println(solve)
	return
}
