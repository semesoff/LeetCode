package main

func setZeroes(matrix [][]int) {
	r := len(matrix)
	c := len(matrix[0])
	rows := make([]bool, r)
	cols := make([]bool, c)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}) // [[1,0,1],[0,0,0],[1,0,1]]
}
