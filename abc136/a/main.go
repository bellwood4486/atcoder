package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	fmt.Println(solve(A, B, C))
}

func solve(A, B, C int) int {
	d := C - (A - B)
	if d < 0 {
		return 0
	} else {
		return d
	}
}
