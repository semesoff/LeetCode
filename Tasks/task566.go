package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	row, col := len(mat), len(mat[0])
	if r*c != row*col {
		return mat
	}
	newMat := make([][]int, r)
	for i := range newMat {
		newMat[i] = make([]int, c)
	}

	cr, cc := 0, 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			newMat[i][j] = mat[cr][cc]
			cc++
			if cc == col {
				cr++
				cc = 0
			}
		}
	}
	return newMat
}

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
}
