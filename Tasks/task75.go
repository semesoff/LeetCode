package main

import "fmt"

func quickSort(array *[]int) []int {
	if len(*array) < 2 {
		return *array
	}

	pivot := (*array)[0]
	var left []int
	var right []int

	for i := 1; i < len(*array); i++ {
		if pivot > (*array)[i] {
			left = append(left, (*array)[i])
		} else {
			right = append(right, (*array)[i])
		}
	}
	return append(append(quickSort(&left), pivot), quickSort(&right)...)
}

func bubbleSort(array *[]int) {
	for {
		flag := 1
		for i := 0; i < len(*array)-1; i++ {
			if (*array)[i] > (*array)[i+1] {
				(*array)[i], (*array)[i+1] = (*array)[i+1], (*array)[i]
				flag = 0
			}
		}
		if flag == 1 {
			break
		}
	}
}

func sortColors(nums []int) {
	bubbleSort(&nums)
}

func main() {
	array := []int{3, 6, 8, 10, 1, 2, 1}
	bubbleSort(&array)
	fmt.Println(array)
}
