package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	stack := make([]int, 0, n)
	for i := 0; i < 2*n; i++ {
		i := i % n
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
