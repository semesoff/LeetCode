package main

import (
	"fmt"
	"sort"
)

type S struct {
	reward1, reward2 []int
}

func (s S) Len() int {
	return len(s.reward1)
}

func (s S) Swap(i, j int) {
	s.reward1[i], s.reward1[j] = s.reward1[j], s.reward1[i]
	s.reward2[i], s.reward2[j] = s.reward2[j], s.reward2[i]
}

func (s S) Less(i, j int) bool {
	return s.reward1[i]-s.reward2[i] > s.reward1[j]-s.reward2[j]
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	sort.Sort(S{reward1, reward2})
	total := 0
	for _, i := range reward1[:k] {
		total += i
	}
	for _, i := range reward2[k:] {
		total += i
	}
	return total
}

func main() {
	fmt.Println(miceAndCheese([]int{1, 1, 3, 4}, []int{4, 4, 1, 1}, 2))
}
