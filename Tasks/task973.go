package main

import (
	"fmt"
	"math"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	distances := make(map[[2]int]float64)
	sliceDistances := make([][]int, 0)
	for _, i := range points {
		distances[[2]int(i)] = math.Sqrt(math.Pow(math.Abs(float64(i[0]-0)), 2) + math.Pow(math.Abs(float64(i[1]-0)), 2))
		sliceDistances = append(sliceDistances, i)
		fmt.Println(i, distances[[2]int(i)])
	}
	sort.Slice(sliceDistances, func(i, j int) bool {
		return distances[[2]int(sliceDistances[i])] <
			distances[[2]int(sliceDistances[j])]
	})
	return sliceDistances[:k]
}

func main() {
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
