package main

import (
	"fmt"
	"slices"
	"sort"
)

func minGroups(intervals [][]int) int {
	startTimes := intervals
	endTimes := slices.Clone(intervals)
	sort.Slice(startTimes, func(i, j int) bool {
		return startTimes[i][0] < startTimes[j][0]
	})
	sort.Slice(endTimes, func(i, j int) bool {
		return endTimes[i][1] < endTimes[j][1]
	})

	endPointer := 0
	countGroups := 0
	for _, start := range startTimes {
		if start[0] > endTimes[endPointer][1] {
			endPointer++
		} else {
			countGroups++
		}
	}
	return countGroups
}

func main() {
	fmt.Println(minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}))
}
