package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func solved(a, b string) string {
	return b + " " + a
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextString()
	b := nextString()
	solve := solved(a, b)
	fmt.Println(solve)
	return
}
