package main

import "fmt"

func uniqueNumbers(grid [][]int) bool {
	hash := make(map[int]int)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if hash[grid[i][j]] == 1 || (grid[i][j] < 1 || grid[i][j] > 9) {
				return false
			}
			hash[grid[i][j]]++
		}
	}
	return true
}

func sliceSum(s []int) int {
	res := 0
	for _, j := range s {
		res += j
	}
	return res
}

func check(grid []int, sum int) bool {
	if sum != sliceSum(grid) {
		return false
	}
	return true
}

func gridIsGood(grid [][]int) bool {
	if !uniqueNumbers(grid) {
		return false
	}
	sum := sliceSum(grid[0])
	if !check([]int{grid[0][0], grid[1][1], grid[2][2]}, sum) {
		return false
	}
	if !check([]int{grid[2][0], grid[1][1], grid[0][2]}, sum) {
		return false
	}
	for i := 1; i < 3; i++ {
		if !check(grid[i], sum) {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		col := []int{grid[0][i], grid[1][i], grid[2][i]}
		if !check(col, sum) {
			return false
		}
	}
	return true
}

func createSquare(grid [][]int, row, col int) [][]int {
	res := make([][]int, 3)
	for i := range res {
		res[i] = make([]int, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i][j] = grid[row+i][col+j]
		}
	}
	return res
}

func numMagicSquaresInside(grid [][]int) int {
	n := len(grid[0])
	m := len(grid)
	total := 0
	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			if gridIsGood(createSquare(grid, i, j)) {
				total++
			}
		}
	}
	return total
}

func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
}
