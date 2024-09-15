package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func cloneTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	// Рекурсивно клонируем левое и правое поддеревья
	return &TreeNode{
		Val:   node.Val,
		Left:  cloneTree(node.Left),
		Right: cloneTree(node.Right),
	}
}

func invertTree(root *TreeNode) *TreeNode {
	var invert func(node *TreeNode) *TreeNode
	invert = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left, node.Right = node.Right, node.Left
		invert(node.Left)
		invert(node.Right)
		return node
	}
	return invert(root)
}

func main() {
	tree := &TreeNode{4, nil, nil}
	tree.Left = &TreeNode{2, nil, nil}
	tree.Right = &TreeNode{7, nil, nil}
	tree.Left.Left = &TreeNode{1, nil, nil}
	tree.Left.Right = &TreeNode{3, nil, nil}
	tree.Right.Right = &TreeNode{9, nil, nil}
	tree.Right.Left = &TreeNode{6, nil, nil}
	fmt.Println(invertTree(tree))
}
