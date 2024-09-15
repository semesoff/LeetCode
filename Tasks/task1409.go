package main

import (
	"fmt"
	"slices"
)

func processQueries(queries []int, m int) []int {
	res := make([]int, len(queries))
	p := make([]int, m)
	for i := 1; i < m+1; i++ {
		p[i-1] = i
	}

	for idx, i := range queries {
		res[idx] = slices.Index(p, i)
		p = slices.DeleteFunc(p, func(j int) bool {
			return j == i
		})
		p = slices.Insert(p, 0, i)
	}
	return res
}

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}
