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

func solved(a string) string {
	b := "0" + a
	return b[0:4]
}

func main() {
	var a string
	fmt.Scan(&a)
	solve := solved(a)
	fmt.Println(solve)
	return
}
