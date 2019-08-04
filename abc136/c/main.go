package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func newSplitScanner(r io.Reader) *bufio.Scanner {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	return sc
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	//
	//	str := `
	//1
	//1000000000
	//`
	//	var sc = newSplitScanner(strings.NewReader(str))

	var sc = newSplitScanner(os.Stdin)
	N := nextInt(sc)
	H := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = nextInt(sc)
	}

	fmt.Println(solve(N, H))
}

func solve(N int, H []int) string {
	for i := len(H) - 1; i > 0; i-- {
		a := H[i-1] - H[i]
		if a >= 2 {
			return "No"
		} else if a == 1 {
			H[i-1]--
		}
	}
	return "Yes"
}
