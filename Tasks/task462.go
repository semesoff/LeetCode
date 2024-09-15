package main

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	mid, res := nums[n/2], 0
	for _, num := range nums {
		res += int(math.Abs(float64(num - mid)))
	}
	return res
}

func main() {

}
