package main

import "fmt"

func subarraysWithKDistinct(nums []int, k int) int {
	n := len(nums)
	countSubArrays := 0
	left, right := 0, 0
	uniqueCount := 0
	frequency := make(map[int]int)

	for right < n {
		// Add current element in the window
		if frequency[nums[right]] == 0 {
			uniqueCount++
		}
		frequency[nums[right]]++
		right++

		// Reducing left, if count of unique elements > k
		for uniqueCount > k {
			frequency[nums[left]]--
			if frequency[nums[left]] == 0 {
				uniqueCount--
			}
			left++
		}

		// If window count == k unique elements, count subArrays
		if uniqueCount == k {
			tempLeft := left
			for uniqueCount == k {
				countSubArrays++
				frequency[nums[tempLeft]]--
				if frequency[nums[tempLeft]] == 0 {
					uniqueCount--
				}
				tempLeft++
			}

			// Restoring the state of frequency and uniqueCount
			for i := left; i < tempLeft; i++ {
				frequency[nums[i]]++
				if frequency[nums[i]] == 1 {
					uniqueCount++
				}
			}
		}
	}
	return countSubArrays
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
}
