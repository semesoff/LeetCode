package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	answers := make([]string, n)
	sortedScore := make([]int, n)
	copy(sortedScore, score)
	sort.Slice(sortedScore, func(i, j int) bool {
		return sortedScore[i] > sortedScore[j]
	})
	for i, j := range score {
		switch index := slices.Index(sortedScore, j) + 1; index {
		case 1:
			answers[i] = "Gold Medal"
		case 2:
			answers[i] = "Silver Medal"
		case 3:
			answers[i] = "Bronze Medal"
		default:
			answers[i] = strconv.Itoa(index)
		}
	}
	return answers
}

func main() {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
