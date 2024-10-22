package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	sort.Ints(nums)
	maxLen := 0
	previous := nums[0]
	curLen := 1
	for i := 1; i < n; i++ {
		if !(previous == nums[i]-1) || previous == nums[i] {
			if previous == nums[i] {
				continue
			}
			maxLen = max(maxLen, curLen)
			previous = nums[i]
			curLen = 1
		} else if previous == nums[i] {

		} else {
			previous = nums[i]
			curLen++
		}
	}
	maxLen = max(maxLen, curLen)
	return maxLen
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 1, 1, 2}))
}
