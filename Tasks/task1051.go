package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	result := 0

	for i := 0; i < len(expected); i++ {
		if expected[i] != heights[i] {
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
