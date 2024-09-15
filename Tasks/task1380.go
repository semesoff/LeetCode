package main

import (
	"fmt"
	"math"
)

func minNumber(slice []int) (int, int) {
	var index int
	number := int(math.Pow(10, 7))
	for j, i := range slice {
		if i < number {
			number = i
			index = j
		}
	}
	return number, index
}

func maxNumber(slice [][]int, index int) int {
	number := int(math.Pow(-10, 7))
	for _, i := range slice {
		if i[index] > number {
			number = i[index]
		}
	}
	return number
}

func luckyNumbers(matrix [][]int) []int {
	m := len(matrix)
	numbers := make([]int, 0)

	for i := 0; i < m; i++ {
		miNumber, index := minNumber(matrix[i])
		maNumber := maxNumber(matrix, index)
		if miNumber == maNumber {
			numbers = append(numbers, miNumber)
		}
	}
	return numbers
}

func main() {
	fmt.Println(luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}))
}
