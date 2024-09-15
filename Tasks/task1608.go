package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	for x := 1; x < n+1; x++ {
		count := 0

		for _, num := range nums {
			if num >= x {
				count++
			}
		}

		if count == x {
			return x
		}
	}

	return -1
}

func main() {
	value := specialArray([]int{2, 3})
	fmt.Println(value)
}
