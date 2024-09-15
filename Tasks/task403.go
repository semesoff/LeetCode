package main

import "fmt"

type Set []int

func (s *Set) add(value int) {
	for _, i := range *s {
		if i == value {
			return
		}
	}
	*s = append(*s, value)
}

func canCross(stones []int) bool {
	dp := make(map[int]Set)
	for _, stone := range stones {
		dp[stone] = Set{}
	}
	dp[0] = Set{0}

	for _, stone := range stones {
		for _, jump := range dp[stone] {
			for distance := jump - 1; distance <= jump+1; distance++ {
				if _, ok := dp[stone+distance]; distance > 0 && ok {
					set := dp[stone+distance]
					set.add(distance)
					dp[stone+distance] = set
				}
			}
		}
	}
	return len(dp[stones[len(stones)-1]]) > 0
}

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
}
