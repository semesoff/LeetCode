package main

import (
	"fmt"
	"sort"
)

type Row struct {
	Index    int
	Strength int
}

func kWeakestRows(mat [][]int, k int) []int {
	rows := make([]Row, len(mat))
	for i := range mat {
		strength := 0
		for j := range mat[i] {
			strength += mat[i][j]
		}
		rows[i] = Row{i, strength}
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Strength == rows[j].Strength {
			return rows[i].Index < rows[j].Index
		}
		return rows[i].Strength < rows[j].Strength
	})
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = rows[i].Index
	}
	return result
}

func main() {
	fmt.Println(kWeakestRows([][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3))
}
