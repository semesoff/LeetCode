package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(node *TreeNode, isLeft bool) int
	dfs = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			if isLeft {
				return node.Val
			}
			return 0
		}
		res := dfs(node.Left, true)
		res += dfs(node.Right, false)
		return res
	}
	return dfs(root, false)
}

func main() {
	tree := &TreeNode{3, nil, nil}
	tree.Left = &TreeNode{9, nil, nil}
	tree.Right = &TreeNode{20, nil, nil}
	tree.Right.Left = &TreeNode{11, nil, nil}
	fmt.Println(sumOfLeftLeaves(tree))
}
