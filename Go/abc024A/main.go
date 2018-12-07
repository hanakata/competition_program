package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStringToInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func solved(a, b, c, k, s, t int) int {
	ans := s*a + t*b
	if k <= s+t {
		ans = ans - c*(s+t)
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextStringToInt()
	b := nextStringToInt()
	c := nextStringToInt()
	k := nextStringToInt()
	s := nextStringToInt()
	t := nextStringToInt()
	solve := solved(a, b, c, k, s, t)
	fmt.Println(solve)
	return
}
