package main

import (
	"fmt"
	"sort"
)

func toSlice(s map[int]int) []int {
	slice := make([]int, 0, len(s))
	for v := range s {
		slice = append(slice, v)
	}
	sort.Ints(slice)
	return slice
}

func getAncestors(n int, edges [][]int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, 0)
	}
	hash := make(map[int][]int)
	for _, val := range edges {
		hash[val[1]] = append(hash[val[1]], val[0])
	}
	memo := make(map[int]map[int]int)
	var dfs func(node int) map[int]int
	dfs = func(node int) map[int]int {
		if _, ok := memo[node]; ok {
			return memo[node]
		}
		ancestors := make(map[int]int)
		for _, parent := range hash[node] {
			ancestors[parent] = 1
			for i := range dfs(parent) {
				ancestors[i] = 1
			}
		}
		memo[node] = ancestors
		res[node] = toSlice(ancestors)
		return ancestors
	}

	for i := 0; i < n; i++ {
		dfs(i)
	}
	return res
}

func main() {
	fmt.Println(getAncestors(8, [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}))
}
