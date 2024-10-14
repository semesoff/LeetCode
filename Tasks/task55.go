package main

import "fmt"

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return true
}

// Вспомогательная функция для нахождения максимума
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // true
}
