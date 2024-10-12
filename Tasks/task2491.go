package main

import (
	"fmt"
	"sort"
)

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	l, r := 0, len(skill)-1
	res := 0
	lastSum := 0
	for l < r {
		sum := skill[l] + skill[r]
		if l == 0 {
			lastSum = sum
			res += skill[l] * skill[r]
		} else {
			res += skill[l] * skill[r]
			if sum != lastSum {
				return -1
			}
		}
		l++
		r--
	}
	return int64(res)
}

func main() {
	fmt.Println(dividePlayers([]int{3,2,5,1,3,4}))
}