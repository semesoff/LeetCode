package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, nodes)
	*nodes = append(*nodes, root.Val)
	inOrderTraversal(root.Right, nodes)
}

func balanceBST(root *TreeNode) *TreeNode {
	nodes := make([]int, 0)
	inOrderTraversal(root, &nodes)
	return buildBalancedBST(nodes, 0, len(nodes)-1)
}

func buildBalancedBST(nodes []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &TreeNode{Val: nodes[mid]}
	root.Left = buildBalancedBST(nodes, start, mid-1)
	root.Right = buildBalancedBST(nodes, mid+1, end)
	return root
}

func main() {
	tree := &TreeNode{1, nil, nil}
	tree.Right = &TreeNode{2, nil, nil}
	tree.Right.Right = &TreeNode{3, nil, nil}
	tree.Right.Right.Right = &TreeNode{4, nil, nil}
	fmt.Println(balanceBST(tree))
}
