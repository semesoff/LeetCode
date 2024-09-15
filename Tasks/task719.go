package main

import (
	"fmt"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	countPairs := func(mid int) int {
		count, left := 0, 0
		for right := 0; right < n; right++ {
			for nums[right]-nums[left] > mid {
				left++
			}
			count += right - left
		}
		return count
	}

	low, high := 0, nums[n-1]-nums[0]
	for low < high {
		mid := (low + high) / 2
		if countPairs(mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	fmt.Println(smallestDistancePair([]int{1, 6, 1}, 3))
}
