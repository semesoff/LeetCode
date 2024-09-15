package main

import "fmt"

func numEnclaves(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	visited := make([][]int, row)
	for i := range visited {
		visited[i] = make([]int, col)
	}

	var findIsland func(r, c int, area *int, isBorder *bool)
	findIsland = func(r, c int, area *int, isBorder *bool) {
		if grid[r][c] == 0 || visited[r][c] == 1 {
			return
		}
		visited[r][c] = 1
		if r == 0 || r == row-1 || c == 0 || c == col-1 {
			*isBorder = true
		}
		*area++
		for _, i := range [4][2]int{{r, c - 1}, {r, c + 1}, {r - 1, c}, {r + 1, c}} {
			if !(i[0] < 0 || i[0] == row || i[1] < 0 || i[1] == col) {
				findIsland(i[0], i[1], area, isBorder)
			}
		}
	}

	res := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 && visited[r][c] == 0 {
				area := 0
				isBorder := false
				findIsland(r, c, &area, &isBorder)
				if !isBorder {
					res += area
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(numEnclaves([][]int{
		{0, 1, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}))
}
