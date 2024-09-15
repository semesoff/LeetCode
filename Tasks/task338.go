package main

import (
	"fmt"
)

func countBits(n int) []int {
	res := make([]int, n+1)
	if n == 0 {
		return res
	}
	res[0] = 0
	if n >= 1 {
		res[1] = 1
	}

	for i := 2; i < n+1; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1

		}

	}
	return res
}

func main() {
	fmt.Println(countBits(5))
}
