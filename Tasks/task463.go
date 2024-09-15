package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	perimeter := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				if i+1 == r || grid[i+1][j] == 0 {
					perimeter++
				}
				if i-1 == -1 || grid[i-1][j] == 0 {
					perimeter++
				}
				if j+1 == c || grid[i][j+1] == 0 {
					perimeter++
				}
				if j-1 == -1 || grid[i][j-1] == 0 {
					perimeter++
				}
			}
		}
	}
	return perimeter
}

func main() {
	fmt.Println(islandPerimeter([][]int{{1}}))
}
