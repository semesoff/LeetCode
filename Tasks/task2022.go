package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return make([][]int, 0)
	}
	res := make([][]int, m)
	idx := 0
	sub := make([]int, 0)
	k := 0
	for _, val := range original {
		k++
		sub = append(sub, val)
		if k == n {
			res[idx] = sub
			k = 0
			idx++
			sub = make([]int, 0)
		}
	}
	return res
}

func main() {
	fmt.Println(construct2DArray([]int{1, 2}, 1, 1))
}
