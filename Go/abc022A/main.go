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

func solved(j, s, t int) bool {
	ans := false
	if s <= j && j <= t {
		ans = true
	}
	return ans
}

func main() {
	var i, ans, j int
	sc.Split(bufio.ScanWords)
	a := nextString()
	b := nextString()
	c := nextString()
	n, _ := strconv.Atoi(a)
	s, _ := strconv.Atoi(b)
	t, _ := strconv.Atoi(c)
	for i < n {
		bS := nextString()
		bI, _ := strconv.Atoi(bS)
		j += bI
		solve := solved(j, s, t)
		if solve {
			ans++
		}
		i++
	}
	fmt.Println(ans)
	return
}
