package main

import "fmt"

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	maxOr := 0
	for _, num := range nums {
		maxOr |= num
	}
	count := 0
	var backtracking func(int, int)
	backtracking = func(i, or int) {
		if i == n {
			if or == maxOr {
				count++
			}
			return
		}
		backtracking(i+1, or)
		backtracking(i+1, or|nums[i])
	}
	backtracking(0, 0)
	return count
}

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))
}
