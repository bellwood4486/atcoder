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
	var sc = newSplitScanner(os.Stdin)
	n := nextInt(sc)

	fmt.Println(n * n * n)
}
