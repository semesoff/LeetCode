package main

import (
	"fmt"
	"math"
	"slices"
)

func minMax(nums []int) (int, int) {
	nMax := math.MinInt32
	nMin := math.MaxInt32
	for _, i := range nums {
		nMax = max(nMax, i)
		nMin = min(nMin, i)
	}
	return nMin, nMax
}

func minimumDeletions(nums []int) int {
	nMin, nMax := minMax(nums)
	minIndex, maxIndex := slices.Index(nums, nMin), slices.Index(nums, nMax)
	minEl, maxEl := min(minIndex, maxIndex), max(minIndex, maxIndex)
	res1 := maxEl + 1
	res2 := minEl + 1 + len(nums) - maxEl
	res3 := len(nums) - minEl
	return min(res1, res2, res3)
}

func main() {
	fmt.Println(minimumDeletions([]int{111}))
}
