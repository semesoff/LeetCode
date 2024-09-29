package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	index := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	nums = nums[:index]
	return index
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 3, 0, 4, 4}, 3)) // 2
}