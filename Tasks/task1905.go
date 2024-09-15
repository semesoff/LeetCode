package main

import "fmt"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	r, c := len(grid1), len(grid1[0])
	visited2 := make([][]int, r)
	for i := range visited2 {
		visited2[i] = make([]int, c)
	}
	islands2 := make([][][2]int, 0)

	var findIslands func(row, col int, cells *[][2]int, islands *[][][2]int, grid *[][]int, visited *[][]int)
	findIslands = func(row, col int, cells *[][2]int, islands *[][][2]int, grid *[][]int, visited *[][]int) {
		if row < 0 || col < 0 || row >= r || col >= c || (*visited)[row][col] == 1 || (*grid)[row][col] == 0 {
			return
		}
		(*visited)[row][col] = 1
		*cells = append(*cells, [2]int{row, col})
		for _, i := range [][2]int{{row, col - 1}, {row, col + 1}, {row - 1, col}, {row + 1, col}} {
			findIslands(i[0], i[1], cells, islands, grid, visited)
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			cells := make([][2]int, 0)
			if grid2[i][j] == 1 && visited2[i][j] == 0 {
				findIslands(i, j, &cells, &islands2, &grid2, &visited2)
				islands2 = append(islands2, cells)
			}
		}
	}

	res := 0
	for _, j := range islands2 {
		flag := true
		for _, cell := range j {
			if grid1[cell[0]][cell[1]] == 0 {
				flag = false
			}
		}
		if flag {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countSubIslands([][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
		[][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}))
}
