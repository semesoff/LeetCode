package main

import (
	"fmt"
	"math"
)

func maximumProduct(nums []int) int {
	max1, max2, max3, min1, min2 := math.MinInt32, math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
	for _, i := range nums {
		if i > max1 {
			max3 = max2
			max2 = max1
			max1 = i
		} else if i > max2 {
			max3 = max2
			max2 = i
		} else if i > max3 {
			max3 = i
		}

		if i < min1 {
			min2 = min1
			min1 = i
		} else if i < min2 {
			min2 = i
		}
	}
	return int(math.Max(float64(max1*max2*max3), float64(min1*min2*max1)))
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
}
