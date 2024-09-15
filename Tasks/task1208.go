package main

import (
	"fmt"
	"math"
)

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	var costs []int

	for i := 0; i < n; i++ {
		costs = append(costs, int(math.Abs(float64(int(s[i])-int(t[i])))))
	}

	left := 0
	current_cost := 0
	max_length := 0

	for right := 0; right < n; right++ {
		current_cost += costs[right]

		for current_cost > maxCost {
			current_cost -= costs[left]
			left++
		}

		max_length = int(math.Max(float64(max_length), float64(right-left+1)))
	}

	return max_length

}

func main() {
	result := equalSubstring("abcd", "bcdf", 3)
	fmt.Println(result)
}
