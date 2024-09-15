package main

func changeCord(i, j, vector *int, m, n int, matrix [][]int) {
	switch *vector {
	case 1:
		if *i-1 < 0 || matrix[*i-1][*j] != -1 {
			*j++
			*vector = 2
		} else {
			*i--
		}
	case 2:
		if *j+1 == n || matrix[*i][*j+1] != -1 {
			*i++
			*vector = 3
		} else {
			*j++
		}
	case 3:
		if *i+1 == m || matrix[*i+1][*j] != -1 {
			*j--
			*vector = 4
		} else {
			*i++
		}
	case 4:
		if *j-1 < 0 || matrix[*i][*j-1] != -1 {
			*i--
			*vector = 1
		} else {
			*j--
		}
	}
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	vector := 2
	i, j := 0, 0
	for node := 1; node <= n*n; node++ {
		matrix[i][j] = node
		changeCord(&i, &j, &vector, n, n, matrix)
	}
	return matrix
}

func main() {

}
