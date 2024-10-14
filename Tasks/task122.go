package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	profit := 0
	stock := math.MaxInt32
	for _, price := range prices {
		if price < stock {
			stock = price
		}
		if price > stock {
			profit += price - stock
			stock = price
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1})) // 7
}
