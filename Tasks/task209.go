package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	right := 0
	minLen := math.MaxInt32
	sum := 0

	for right < n {
		sum += nums[right]
		for sum >= target {
			minLen = int(math.Min(float64(minLen), float64(right-left+1)))
			sum -= nums[left]
			left++
		}
		right++
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
