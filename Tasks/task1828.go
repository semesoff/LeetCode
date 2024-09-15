package main

import (
	"fmt"
)

func countPoints(points [][]int, queries [][]int) []int {
	ps := make([]int, len(queries))
	for i, circ := range queries {
		c := 0
		for _, p := range points {
			if sq(circ[0]-p[0])+sq(circ[1]-p[1]) <= sq(circ[2]) {
				c++
			}
		}
		ps[i] = c
	}
	return ps
}

func sq(val int) int {
	return val * val
}

func main() {
	fmt.Println(countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
}
