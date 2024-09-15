package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	index := 0
	for _, j := range nums {
		if j != 0 {
			nums[index] = j
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}
