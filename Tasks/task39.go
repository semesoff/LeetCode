package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	n := len(candidates)
	var backtracking func(array []int, index, sum int)
	backtracking = func(array []int, index, sum int) {
		if sum == target {
			arr := make([]int, len(array))
			copy(arr, array)
			res = append(res, arr)
			return
		}
		if index == n || sum > target {
			return
		}
		backtracking(array, index+1, sum)

		array = append(array, candidates[index])
		backtracking(array, index, sum+candidates[index])
	}
	backtracking([]int{}, 0, 0)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
