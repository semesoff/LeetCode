package main

import (
	"fmt"
	"sort"
)

func createDict(dict *map[int]int, arr *[]int) {
	for _, i := range *arr {
		(*dict)[i]++
	}
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dict := make(map[int]int)
	createDict(&dict, &arr1)

	var res []int
	var other []int

	for _, i := range arr2 {
		for dict[i] > 0 {
			res = append(res, i)
			dict[i]--
		}
		delete(dict, i)
	}

	for num, count := range dict {
		for count > 0 {
			other = append(other, num)
			count--
		}
	}

	sort.Ints(other)

	return append(res, other...)
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}

	result := relativeSortArray(arr1, arr2)
	fmt.Println(result)
}
