package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backtracking func(counts, path []int)
	backtracking = func(counts, path []int) {
		if len(counts) == 0 {
			res = append(res, path)
			return
		}
		for i := 0; i < len(counts); i++ {
			newNums := append([]int(nil), counts[:i]...)
			newNums = append(newNums, counts[i+1:]...)
			newPath := append([]int(nil), path...)
			newPath = append(newPath, counts[i])
			backtracking(newNums, newPath)
		}
	}
	backtracking(nums, []int(nil))
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
