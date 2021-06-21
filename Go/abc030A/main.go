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

func solved(a float64, b float64, c float64, d float64) string {
	takahashi := b / a
	aoki := d / c
	if takahashi > aoki {
		return "TAKAHASHI"
	}
	if takahashi < aoki {
		return "AOKI"
	}
	return "DRAW"
}

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	solve := solved(a, b, c, d)
	fmt.Println(solve)
	return
}
