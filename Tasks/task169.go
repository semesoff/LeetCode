package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	fmt.Println(majorityElement([]int{1, 1, 1, 1, 1, 4, 5, 2, 2, 9}))
}

// 1 1 1 2 3 4 5 6 7 8
