package main

//func rotate(matrix [][]int) {
//	n := len(matrix)
//	other := make([][]int, n)
//	for i := range other {
//		other[i] = make([]int, n)
//	}
//	for i := range matrix {
//		for j := range matrix[i] {
//			other[j][n-1-i] = matrix[i][j]
//		}
//	}
//	for i := range matrix {
//		for j := range matrix[i] {
//			matrix[i][j] = other[i][j]
//		}
//	}
//}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
