package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	left, right := 0, 0
	oddCount := 0
	count := 0
	prefix := 0

	for right < len(nums) {
		if nums[right]%2 != 0 {
			oddCount++
			prefix = 0
		}

		for oddCount == k {
			prefix++
			if nums[left]%2 != 0 {
				oddCount--
			}
			left++
		}

		count += prefix
		right++
	}

	return count
}

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2}, 2))
}
