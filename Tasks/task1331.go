package main

import (
	"fmt"
	"slices"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	other := slices.Clone(arr)
	sort.Ints(other)

	hash := make(map[int]int)
	k := 1
	for i := 0; i < len(other); i++ {
		if i > 0 && other[i-1] != other[i] {
			k++
		}
		hash[other[i]] = k
	}

	res := make([]int, len(arr))
	for i, j := range arr {
		res[i] = hash[j]
	}

	return res
}

func main() {
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}
