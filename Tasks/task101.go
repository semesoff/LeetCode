package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var sym func(root1 *TreeNode, root2 *TreeNode) bool
	sym = func(root1 *TreeNode, root2 *TreeNode) bool {
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil {
			return false
		}
		if root1.Val != root2.Val {
			return false
		}

		return sym(root1.Left, root2.Right) && sym(root1.Right, root2.Left)
	}
	return sym(root, root)
}

func main() {
	tree := &TreeNode{1, nil, nil}
	tree.Left = &TreeNode{2, nil, nil}
	tree.Right = &TreeNode{2, nil, nil}
	tree.Left.Left = &TreeNode{3, nil, nil}
	tree.Left.Right = &TreeNode{4, nil, nil}
	tree.Right.Left = &TreeNode{4, nil, nil}
	tree.Right.Right = &TreeNode{3, nil, nil}
	fmt.Println(isSymmetric(tree))
}
