package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			if i-m[nums[i]] <= k {
				return true
			}
		}
		m[nums[i]] = i
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
}
