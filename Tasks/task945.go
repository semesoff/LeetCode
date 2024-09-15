package main

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	count := 0

	for i := 1; i < len(nums); i++ {
		previous := nums[i-1]
		current := nums[i]

		if current <= previous {
			current = previous + 1
			count += current - nums[i]
			nums[i] = current
		}
	}
	return count
}
