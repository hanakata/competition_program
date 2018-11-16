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

func solved(n string) string {
	solve := n + "pp"
	return solve
}

func main() {
	n := nextString()
	solve := solved(n)
	fmt.Println(solve)
	return
}
