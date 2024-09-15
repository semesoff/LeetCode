package main

import "fmt"

func stoneGameII(piles []int) int {
	n := len(piles)
	suffixSum := make([]int, n)
	suffixSum[n-1] = piles[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = piles[i] + suffixSum[i+1]
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}

	var dp func(i, m int) int
	dp = func(i, m int) int {
		// if we are in the last segment, take all the stones
		if i == n {
			return 0
		}
		// if this is last step, then we take all remaining stones
		if 2*m >= n-i {
			return suffixSum[i]
		}
		// if the value has already been calculated, we return it
		if memo[i][m] > 0 {
			return memo[i][m]
		}
		// calculating the maximum number of stones that the current player can take
		maxStones := 0
		for x := 1; x <= 2*m && i+x <= n; x++ {
			maxStones = max(maxStones, suffixSum[i]-dp(i+x, max(m, x)))
		}
		memo[i][m] = maxStones
		return maxStones
	}
	return dp(0, 1)
}

func main() {
	fmt.Println(stoneGameII([]int{2, 7, 9, 4, 4}))
}
