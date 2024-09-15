package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {
	n := len(fruits)
	hash := map[int]int{}

	left := 0
	maxTotal := 1
	hash[fruits[left]]++

	for right := 1; right < n; right++ {
		hash[fruits[right]]++
		for len(hash) > 2 {
			hash[fruits[left]]--
			if hash[fruits[left]] == 0 {
				delete(hash, fruits[left])
			}
			left++
		}
		if right-left+1 > maxTotal {
			maxTotal = right - left + 1
		}
	}
	return maxTotal
}

func main() {
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
}
