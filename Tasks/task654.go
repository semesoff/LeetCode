package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxVal(nums []int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}
	maxn, idx := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxn {
			maxn = nums[i]
			idx = i
		}
	}
	return maxn, idx
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxn, idx := maxVal(nums)
	root := &TreeNode{
		Val:   maxn,
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}
	return root
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{3, 2, 1}))
}
