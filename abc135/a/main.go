package main

import (
	"fmt"
	"math"
	"strconv"
)

func solve(a, b int) string {
	m := math.Abs(float64((a - b) % 2))
	if m > 0 {
		return "IMPOSSIBLE"

	}
	ans := ((a - b) / 2) + b
	return strconv.Itoa(ans)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(solve(a, b))
}
