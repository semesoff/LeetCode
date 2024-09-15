package main

import "fmt"

func numIslands(grid [][]byte) int {
	r, c := len(grid), len(grid[0])
	visited := make(map[[2]int]int)
	countIslands := 0
	var findIslands func(i, j int)
	findIslands = func(i, j int) {
		if i < 0 || i >= r || j < 0 || j >= c {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		visited[[2]int{i, j}] = 1
		for _, k := range [][2]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
			if _, ok := visited[[2]int{k[0], k[1]}]; !ok {
				findIslands(k[0], k[1])
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if _, ok := visited[[2]int{i, j}]; !ok && grid[i][j] == '1' {
				findIslands(i, j)
				countIslands++
			}
		}
	}
	return countIslands
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))

}
