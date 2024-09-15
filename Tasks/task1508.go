package main

import (
	"fmt"
	"sort"
)

const MOD = 1e9 + 7

func sliceSum(s []int) int {
	res := 0
	for _, j := range s {
		res += j
		res %= MOD
	}
	return res
}

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0)
	for i := 0; i < n; i++ {
		curSum := nums[i]
		sums = append(sums, curSum)
		for j := i + 1; j < n; j++ {
			curSum += nums[j]
			sums = append(sums, curSum)
		}
	}
	sort.Ints(sums)
	return sliceSum(sums[left-1 : right])
}

func main() {
	fmt.Println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
}
