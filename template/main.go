package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var sc = newSplitScanner(os.Stdin)

func newScanner(r io.Reader) *bufio.Scanner {
	return bufio.NewScanner(r)
}

func newSplitScanner(r io.Reader) *bufio.Scanner {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	return sc
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	fmt.Println("hello")
}
