package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStringToInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func solved(n int) int {
	var x int
	x_a := n % 2
	if x_a == 0 {
		x = n / 2
	} else {
		var x_b float64
		x_b = float64(n) / 2
		x = int(math.Ceil(x_b))
	}
	ans := (n - x) * x
	return ans
}

func main() {
	a := nextStringToInt()
	solve := solved(a)
	fmt.Println(solve)
	return
}
