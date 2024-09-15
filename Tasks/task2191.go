package main

import (
	"fmt"
	"sort"
	"strconv"
)

func otherSlice(otherNums *[]int, mapping []int, nums []int) {
	for _, num := range nums {
		current := ""
		for _, charNum := range strconv.Itoa(num) {
			cur, _ := strconv.Atoi(string(charNum))
			current += strconv.Itoa(mapping[cur])
		}
		resNum, _ := strconv.Atoi(current)
		*otherNums = append(*otherNums, resNum)
	}
}

func sortJumbled(mapping []int, nums []int) []int {
	type Pair struct {
		Original    int
		Transformed int
	}
	otherNums := make([]int, 0)
	otherSlice(&otherNums, mapping, nums)

	pairs := make([]Pair, len(nums))
	for i, num := range nums {
		pairs[i] = Pair{num, otherNums[i]}
	}
	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].Transformed == pairs[j].Transformed {
			return false
		}
		return pairs[i].Transformed < pairs[j].Transformed
	})
	for i, pair := range pairs {
		nums[i] = pair.Original
	}
	return nums
}

func main() {
	mapping := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
	nums := []int{991, 338, 38}
	fmt.Println(sortJumbled(mapping, nums))
}
