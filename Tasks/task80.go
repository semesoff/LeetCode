package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	t := n
	deletedCount := 0
	k := 1
	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			k++
		} else {
			k = 1
		}
		if k > 2 {
			deletedCount++
			nums = append(nums[:i], nums[i+1:]...)
			n--
			i--
		}
	}

	return t - deletedCount
}

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,1,2,3,3}))
}