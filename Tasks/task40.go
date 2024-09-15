package main

import (
	"fmt"
	"slices"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var backtrack func(start, sum int, path []int)

	backtrack = func(start, sum int, path []int) {
		if sum == target {
			result = append(result, slices.Clone(path))
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if sum+candidates[i] > target {
				break
			}
			backtrack(i+1, sum+candidates[i], append(path, candidates[i]))
		}
	}
	backtrack(0, 0, []int{})
	return result
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
