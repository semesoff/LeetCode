package main

import (
	"fmt"
)

func prisonAfterNDays(cells []int, n int) []int {
	next := func(x []int) []int {
		r := make([]int, len(x))
		for i := 1; i < len(x)-1; i++ {
			if x[i-1] == x[i+1] {
				r[i] = 1
			} else {
				r[i] = 0
			}
		}
		return r
	}
	cmp := func(a, b []int) bool {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	cells = next(cells)
	n--
	nds := append([]int{}, cells...)
	for i := 0; i < n; i++ {
		nds = next(nds)
		if cmp(nds, cells) {
			n = n % (i + 1)
			i = -1
		}
	}
	return nds
}

func main() {
	fmt.Println(prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000))
}
