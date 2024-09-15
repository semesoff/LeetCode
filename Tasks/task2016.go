package main

import (
	"fmt"
	"math"
)

func maximumDifference(nums []int) int {
	n := len(nums)
	minValue := nums[0]
	maxDiff := -1

	for i := 1; i < n; i++ {
		if nums[i] > minValue {
			maxDiff = int(math.Max(float64(maxDiff), float64(nums[i]-minValue)))
		} else {
			minValue = nums[i]
		}
	}

	return maxDiff
}
func main() {
	fmt.Println(maximumDifference([]int{87, 68, 91, 86, 58, 63, 43, 98, 6, 40}))
}
