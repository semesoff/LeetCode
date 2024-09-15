package main

func maxPoints(points [][]int) int64 {
	m, n := len(points), len(points[0])
	dp := make([]int, n)
	copy(dp, points[0])

	for i := 1; i < m; i++ {
		left := make([]int, n)
		right := make([]int, n)

		left[0] = dp[0]
		for j := 1; j < n; j++ {
			left[j] = max(left[j-1]-1, dp[j])
		}

		right[n-1] = dp[n-1]
		for j := n - 2; j >= 0; j-- {
			right[j] = max(right[j+1]-1, dp[j])
		}

		for j := 0; j < n; j++ {
			dp[j] = points[i][j] + max(left[j], right[j])
		}
	}

	result := dp[0]
	for _, v := range dp {
		result = max(result, v)
	}

	return int64(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
