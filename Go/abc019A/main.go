package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func solved(a, b, c int) int {
	//sliceじゃないとソートできなかった
	var r = []int{a, b, c}
	sort.Sort(sort.IntSlice(r))
	return r[1]
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextString()
	b := nextString()
	c := nextString()
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	l, _ := strconv.Atoi(c)
	solve := solved(n, m, l)
	fmt.Println(solve)
	return
}
