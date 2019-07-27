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

func isAsc(ps []int) bool {
	for i := 0; i < len(ps)-1; i++ {
		if ps[i] > ps[i+1] {
			return false
		}
	}
	return true
}

func solve(n int, ps []int) string {
	var m [][]int

	for i := 0; i < len(ps)-1; i++ {
		if ps[i] > ps[i+1] {
			m = append(m, []int{i, i + 1})
		}
	}

	if len(m) == 0 {
		return "YES"
	}

	for _, v := range m {
		for _, v2 := range m {
			ps2 := make([]int, len(ps))
			copy(ps2, ps)

			ps2[v[0]], ps2[v2[1]] = ps2[v2[1]], ps2[v[0]]

			if isAsc(ps2) {
				return "YES"
			}
		}
	}

	return "NO"
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	ps := make([]int, n)
	for i := 0; i < n; i++ {
		ps[i] = nextInt()
	}

	fmt.Println(solve(n, ps))
}
