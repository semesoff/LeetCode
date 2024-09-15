package main

import (
	"fmt"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	num := 1
	for i := 0; i < len(nums); i++ {
		if nums[i]-num > 1 {
			return num
		} else {
			num = nums[i]
		}
	}
	return num
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}
