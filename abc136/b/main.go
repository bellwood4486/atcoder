package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	//fmt.Println(solve3(N))
	fmt.Println(answer(N))
}

func answer(N int) int {
	ans := 0
	for i := 1; i <= N; i++ {
		if len(strconv.Itoa(i))%2 == 1 {
			ans++
		}
	}
	return ans
}

func isOddDigit(num int) bool {
	digit := 0
	for num != 0 {
		num /= 10
		digit++
	}
	return digit%2 != 0
}
func solve3(N int) int {
	ans := 0
	for i := 1; i <= N; i++ {
		if isOddDigit(i) {
			ans++
		}
	}
	return ans
}
