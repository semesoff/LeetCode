package main

import "fmt"

func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	index := 0
	for k >= 0 {
		if index == n {
			index = 0
		}
		if k < chalk[index] {
			return index
		}
		k -= chalk[index]
		index++
	}
	return index
}

func main() {
	fmt.Println(chalkReplacer([]int{3, 4, 1, 2}, 25))
}
