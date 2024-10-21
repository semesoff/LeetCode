package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var bst func(node *TreeNode, min, max int) bool
	bst = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return bst(node.Left, min, node.Val) && bst(node.Right, node.Val, max)
	}
	return bst(root, math.MinInt64, math.MaxInt64)
}

func main() {
	tree := &TreeNode{5, nil, nil}
	tree.Left = &TreeNode{4, nil, nil}
	tree.Right = &TreeNode{6, nil, nil}
	tree.Right.Left = &TreeNode{3, nil, nil}
	tree.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(isValidBST(tree))
}
