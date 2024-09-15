package main

import (
	"fmt"
	"math"
)

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	memo := make(map[[2]int]int)

	var recursion func(index, w int) int
	recursion = func(index, w int) int {
		if index == n {
			return 0
		}

		if val, ok := memo[[2]int{index, w}]; ok {
			return val
		}

		cur_w := 0
		cur_h := 0
		res := math.MaxInt32

		for i := index; i < n; i++ {
			if cur_w+books[i][0] > shelfWidth {
				break
			}
			cur_w += books[i][0]
			cur_h = int(math.Max(float64(cur_h), float64(books[i][1])))
			res = int(math.Min(float64(res), float64(cur_h+recursion(i+1, cur_w))))
		}

		memo[[2]int{index, w}] = res
		return res
	}

	return recursion(0, 0)
}

func main() {
	books := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
	shelfWidth := 4
	fmt.Println(minHeightShelves(books, shelfWidth)) // Output: 6
}
