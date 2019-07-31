package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func newScanner(r io.Reader) *bufio.Scanner {
	return bufio.NewScanner(r)
}

func nextLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

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
	str := `
3
1 2 3
`
	var sc = newSplitScanner(strings.NewReader(str))

	//var sc = newSplitScanner(os.Stdin)
	n := nextInt(sc)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt(sc)
	}

	fmt.Println(n, a)
}
