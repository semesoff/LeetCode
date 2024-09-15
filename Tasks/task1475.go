package main

import "fmt"

func finalPrices(prices []int) []int {
	result := make([]int, 0, len(prices))
	for i := 0; i < len(prices); i++ {
		isAdd := false
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				result = append(result, prices[i]-prices[j])
				isAdd = true
				break
			}
		}
		if !isAdd {
			result = append(result, prices[i])
		}
	}
	return result
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}
