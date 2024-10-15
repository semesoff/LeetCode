package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	mLeft := 0
	for i := 0; i < n; i++ {
		maxLeft[i] = mLeft
		if height[i] > mLeft {
			mLeft = height[i]
		}
	}

	mRight := 0
	for i := n - 1; i >= 0; i-- {
		maxRight[i] = mRight
		if height[i] > mRight {
			mRight = height[i]
		}
	}

	res := 0
	for i := range height {
		if val := min(maxLeft[i], maxRight[i]) - height[i]; val > 0 {
			res += val
		}
	}
	return res
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
}
