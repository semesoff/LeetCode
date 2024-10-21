package main

import "fmt"

func maximumSwap(num int) int {
	nums := make([]int, 0)
	for num > 0 {
		nums = append([]int{num % 10}, nums...)
		num /= 10
	}

	n := len(nums)
	for i := 0; i < len(nums); i++ {
		index := i
		for j := i + 1; j < n; j++ {
			if nums[j] >= nums[index] {
				index = j
			}
		}
		if index != i && nums[index] != nums[i] {
			nums[i], nums[index] = nums[index], nums[i]
			break
		}
	}
	res := 0
	for _, v := range nums {
		res = res*10 + v
	}
	return res
}

func main() {
	fmt.Println(maximumSwap(2736)) // 7236
}
