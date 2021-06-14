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

func solved(a int) string {
	if a == 100 {
		return "Perfect"
	}
	if 90 <= a && a <= 99 {
		return "Great"
	}
	if 60 <= a && a <= 89 {
		return "Good"
	}
	return "Bad"

}

func main() {
	a := nextStringToInt()
	solve := solved(a)
	fmt.Println(solve)
	return
}
