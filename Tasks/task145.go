package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var traversal func(root *TreeNode, res *[]int) bool
	traversal = func(root *TreeNode, res *[]int) bool {
		if root == nil {
			return true
		}
		ok1 := traversal(root.Left, res)
		ok2 := traversal(root.Right, res)
		if ok1 && ok2 {
			*res = append(*res, root.Val)
			return true
		}
		return false
	}
	arr := make([]int, 0)
	traversal(root, &arr)
	return arr
}

func main() {
	tree := &TreeNode{1, nil, nil}
	tree.Right = &TreeNode{2, nil, nil}
	tree.Right.Left = &TreeNode{3, nil, nil}
	fmt.Println(postorderTraversal(tree))
}
