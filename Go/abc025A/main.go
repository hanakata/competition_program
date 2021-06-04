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

func nextStringToInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func solved(s string, n int) string {
	s_arr := []string{}
	for _, c := range s {
		for _, d := range s {
			s_arr = append(s_arr, string(c)+string(d))
		}
	}
	return s_arr[n-1]
}

func main() {
	s := nextString()
	n := nextStringToInt()
	solve := solved(s, n)
	fmt.Println(solve)
	return
}
