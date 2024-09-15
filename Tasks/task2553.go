package main

import (
	"fmt"
	"slices"
)

func numToSlice(num int) []int {
	res := make([]int, 0)
	for num > 0 {
		item := num % 10
		num /= 10
		res = slices.Insert(res, 0, item)
	}
	return res
}

func separateDigits(nums []int) []int {
	res := make([]int, 0)
	for _, i := range nums {
		res = append(res, numToSlice(i)...)
	}
	return res
}

func main() {
	fmt.Println(separateDigits([]int{123, 456}))
}
