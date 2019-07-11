package main

import "fmt"

// 参考：http://drken1215.hatenablog.com/entry/2019/07/07/233700
func answer(L, R int) int {
	if R-L > 2019 {
		return 0
	}

	r := 2018
	for i := L; i <= R-1; i++ {
		for j := L + 1; j <= R; j++ {
			a := i * j % 2019
			if a < r {
				r = a
			}
		}
	}
	return r
}

func solve(L, R int) int {
	r := 3000
	for i := L; i <= R-1; i++ {
		for j := i + 1; j <= R; j++ {
			a := i * j % 2019
			if a < r {
				r = a
			}
		}
	}
	return r
}

func main() {
	var L, R int
	fmt.Scan(&L, &R)
	//fmt.Println(solve(L, R))
	fmt.Println(answer(L, R))
}
