package main

//func searchMatrix(matrix [][]int, target int) bool {
//	r := len(matrix)
//	c := len(matrix[0])
//	for i := 0; i < r; i++ {
//		if target >= matrix[i][0] && target <= matrix[i][c-1] {
//			for _, v := range matrix[i] {
//				if v == target {
//					return true
//				}
//			}
//		}
//	}
//	return false
//}

func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)
	c := len(matrix[0])
	i := 0
	j := c - 1
	for i < r && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
