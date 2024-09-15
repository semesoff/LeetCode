package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	r, c := len(matrix), len(matrix[0])
	visited := make([][]bool, r)
	for i := range visited {
		visited[i] = make([]bool, c)
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i, j, dir := 0, 0, 0

	for k := 0; k < r*c; k++ {
		res = append(res, matrix[i][j])
		visited[i][j] = true

		ni, nj := i+directions[dir][0], j+directions[dir][1]

		if ni < 0 || ni >= r || nj < 0 || nj >= c || visited[ni][nj] {
			dir = (dir + 1) % 4
			ni, nj = i+directions[dir][0], j+directions[dir][1]
		}

		i, j = ni, nj
	}

	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
