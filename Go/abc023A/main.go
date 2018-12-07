package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func solved(a, b string) int {
	n, _ := strconv.Atoi(a)
	s, _ := strconv.Atoi(b)
	return n + s
}

func main() {
	a := nextString()
	b := strings.Split(a, "")
	solve := solved(b[0], b[1])
	fmt.Println(solve)
	return
}
