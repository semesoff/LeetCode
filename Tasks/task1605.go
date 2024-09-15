package main

import (
	"fmt"
	"math"
)

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}
	i, j := 0, 0

	for i < m && j < n {
		val := int(math.Min(float64(rowSum[i]), float64(colSum[j])))
		result[i][j] = val
		rowSum[i] -= val
		colSum[j] -= val

		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}
	return result
}

func main() {
	fmt.Println(restoreMatrix([]int{3, 8}, []int{4, 7}))
}
