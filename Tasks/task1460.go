package main

import (
	"fmt"
)

func canBeEqual(target []int, arr []int) bool {
	counts := make(map[int]int)

	for i, j := range target {
		counts[arr[i]] = 1
		counts[j] = 1
	}
	return len(counts) == len(target)
}

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}
