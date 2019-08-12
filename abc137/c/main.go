package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func sortString(in string) string {
	out := strings.Split(in, "")
	sort.Strings(out)
	return strings.Join(out, "")
}

func solve(N int, s []string) uint64 {
	for _, v := range s {
		ss := sortString(v)
		cache[ss]++
	}
	return count(cache)
}

func count(cache map[string]int) uint64 {
	ret := uint64(0)
	for _, v := range cache {
		if v >= 2 {
			for i := 1; i <= v-1; i++ {
				ret += uint64(i)
			}
		}
	}
	return ret
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
	var sc = newSplitScanner(os.Stdin)
	n := nextInt(sc)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = nextLine(sc)
	}

	fmt.Println(solve(n, a))
}
