package main

import (
	"fmt"
	"slices"
)

func createTargetArray(nums []int, index []int) []int {
	n := len(nums)
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		res = slices.Insert(res, index[i], nums[i])
		fmt.Println(res)
	}
	return res
}

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}
