package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	profit := 0
	costBuy := math.MaxInt32

	for _, price := range prices {
		if price < costBuy {
			costBuy = price
		} else if price-costBuy > profit {
			profit = price - costBuy
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
}