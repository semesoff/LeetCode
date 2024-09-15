package main

import (
	"fmt"
)

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func quickSort(nums []int, low, high int) {
	if low < high {
		p := partition(nums, low, high)
		quickSort(nums, low, p-1)
		quickSort(nums, p+1, high)
	}
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func main() {
	counts := []int{1, 2, 3, 5, 5, 5, 5}
	fmt.Println(sortArray(counts))
}
