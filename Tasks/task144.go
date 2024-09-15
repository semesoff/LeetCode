package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var traversal func(node *TreeNode, res *[]int) []int
	traversal = func(node *TreeNode, res *[]int) []int {
		if node == nil {
			return *res
		}
		*res = append(*res, node.Val)
		traversal(node.Left, res)
		traversal(node.Right, res)
		return *res
	}
	return traversal(root, &[]int{})
}

func main() {
	tree := &TreeNode{1, nil, nil}
	tree.Right = &TreeNode{2, nil, nil}
	tree.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(preorderTraversal(tree))
}
