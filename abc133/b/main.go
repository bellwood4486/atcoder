package main

import (
	"fmt"
	"math"
)

func solve(n int, d int, xs [][]float64) int {
	var r int
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			var p float64
			for k := 0; k < d; k++ {
				a := math.Abs(xs[i][k] - xs[j][k])
				p += math.Pow(a, 2)
			}
			s := math.Sqrt(p)
			si := int(s)
			if s-float64(si) == 0 {
				r++
			}

		}
	}
	return r
}

func main() {
	var n, d int
	fmt.Scan(&n)
	fmt.Scan(&d)

	var xs [][]float64
	for i := 0; i < n; i++ {
		var x []float64
		for j := 0; j < d; j++ {
			var item float64
			fmt.Scan(&item)
			x = append(x, item)
		}
		xs = append(xs, x)
	}

	fmt.Println(solve(n, d, xs))
}
