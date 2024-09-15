package main

import (
	"fmt"
	"sort"
)

func minimumPushes(word string) int {
	frequency := make([]int, 26)
	for _, c := range word {
		frequency[c-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(frequency)))

	total := 0
	for i, freq := range frequency {
		if freq == 0 {
			break
		}
		total += freq * (i/8 + 1)
	}
	return total
}

func main() {
	fmt.Println(minimumPushes("abzaqsqcyrbzsrvamylmyxdjl"))
}
