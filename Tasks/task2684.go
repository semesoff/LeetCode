package main

import "fmt"

func maxMoves(grid [][]int) int {
	rLimit, cLimit := len(grid), len(grid[0])
	memo := make([][]int, rLimit)
	for i := range memo {
		memo[i] = make([]int, cLimit)
	}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if memo[row][col] != 0 {
			return memo[row][col]
		}

		maxLen := 0
		directions := [3][2]int{{-1, 1}, {0, 1}, {1, 1}}
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < rLimit && newCol >= 0 && newCol < cLimit && grid[newRow][newCol] > grid[row][col] {
				maxLen = max(maxLen, 1+dfs(newRow, newCol))
			}
		}
		memo[row][col] = maxLen
		return maxLen
	}

	maxLength := 0
	for row := 0; row < rLimit; row++ {
		maxLength = max(maxLength, dfs(row, 0))
	}
	return maxLength
}

func main() {
	fmt.Println(maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
}
