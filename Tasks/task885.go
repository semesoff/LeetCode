package main

import "fmt"

func movement(vector, row, col *int) {
	switch *vector {
	case 0: // up
		*row--
	case 1: // right
		*col++
	case 2: // down
		*row++
	case 3: // left
		*col--
	}
}

func changeVector(vector *int) {
	switch *vector {
	case 0: // up
		*vector = 1
	case 1: // right
		*vector = 2
	case 2: // down
		*vector = 3
	case 3: // left
		*vector = 0
	}
}

func borderIsGood(cRow, cCol, rows, cols int) bool {
	return (cRow >= 0 && cRow <= rows-1) && (cCol >= 0 && cCol <= cols-1)
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	matrix := make([][]int, rows*cols)
	matrix[0] = []int{rStart, cStart}
	total := rows * cols

	visited := 1
	travelLength := 1
	vector := 1
	k := 1
	freq := 0
	for visited != total {
		freq++
		for i := 0; i < travelLength; i++ {
			movement(&vector, &rStart, &cStart)
			if borderIsGood(rStart, cStart, rows, cols) {
				visited++
				matrix[k] = []int{rStart, cStart}
				k++
				if visited == total {
					return matrix
				}
			}
		}
		changeVector(&vector)
		if freq == 2 {
			freq = 0
			travelLength++
		}
	}

	return matrix
}

func main() {
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}
