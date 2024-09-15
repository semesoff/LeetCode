package main

import (
	"fmt"
	"math"
)

func findSmallest(arrays [][]int, index int) (int, int) {
	minNumber := math.MaxInt32
	indexArray := 0
	for idx, i := range arrays {
		if idx != index && i[0] < minNumber {
			minNumber = i[0]
			indexArray = idx
		}
	}
	return minNumber, indexArray
}

func findLargest(arrays [][]int, index int) (int, int) {
	maxNumber := math.MinInt32
	indexArray := 0
	for idx, i := range arrays {
		if idx != index && i[len(i)-1] > maxNumber {
			maxNumber = i[len(i)-1]
			indexArray = idx
		}
	}
	return maxNumber, indexArray
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func maxDistance(arrays [][]int) int {
	smallest, indexSmallest := findSmallest(arrays, -1)
	largest, _ := findLargest(arrays, indexSmallest)
	largest1, indexLargest1 := findLargest(arrays, -1)
	smallest1, _ := findSmallest(arrays, indexLargest1)
	return max(abs(largest-smallest), abs(largest1-smallest1))
}

func main() {
	fmt.Println(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}))
}
