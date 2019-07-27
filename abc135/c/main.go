package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func solve(n int, as []int, bs []int) int {
	sum := 0
	for i := len(bs) - 1; i >= 0; i-- {
		if as[i+1] >= bs[i] {
			sum += bs[i]
			as[i+1], bs[i] = as[i+1] - bs[i], 0
		} else {
			sum += as[i+1]
			as[i+1], bs[i] = 0, bs[i] - as[i+1]
			if as[i] >= bs[i] {
				sum += bs[i]
				as[i] = as[i] - bs[i]
			} else {
				sum += as[i]
				as[i] = 0
			}
		}
	}
	return sum
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	as := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		as[i] = nextInt()
	}
	bs := make([]int, n)
	for i := 0; i < n; i++ {
		bs[i] = nextInt()
	}

	fmt.Println(solve(n, as, bs))
}
