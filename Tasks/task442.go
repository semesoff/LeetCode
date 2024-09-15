package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	output := []int{}
	for _, num := range nums {
		idx := abs(num)
		if nums[idx-1] < 0 {
			output = append(output, idx)
		} else {
			nums[idx-1] *= -1
		}
	}
	return output
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	fmt.Println(findDuplicates([]int{1, 2, 3, 4, 4}))
}
