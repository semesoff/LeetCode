package main

import (
	"fmt"
)

func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Filling DP table
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1 // A single character needs 1 turn
		for j := i + 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + 1 // Initial assumption, print s[j] separately
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j-1])
				}
			}
		}
	}

	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(strangePrinter("aba"))
}
