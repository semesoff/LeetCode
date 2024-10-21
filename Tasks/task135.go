package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	count := 1
	candies := make([]int, n)
	candies[0] = 1

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			count += (candies[i-1] + 1) - candies[i]
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
			count++
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			count += (candies[i+1] + 1) - candies[i]
			candies[i] = candies[i+1] + 1
		}
	}
	return count
}

func main() {
	fmt.Println(candy([]int{1, 0, 2})) // 5
}
