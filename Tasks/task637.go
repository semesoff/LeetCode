package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	currLevel := []*TreeNode{root}
	for len(currLevel) > 0 {}
}

func main() {
	tree := &TreeNode{3, nil, nil}
	tree.Left = &TreeNode{9, nil, nil}
	tree.Right = &TreeNode{20, nil, nil}
	tree.Right.Left = &TreeNode{15, nil, nil}
	tree.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(averageOfLevels(tree))
}
