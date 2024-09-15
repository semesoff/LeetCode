package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var sameTree func(node1 *TreeNode, node2 *TreeNode) bool
	sameTree = func(node1 *TreeNode, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil {
			return false
		}

		res1 := sameTree(node1.Left, node2.Left)
		res2 := sameTree(node1.Right, node2.Right)

		return node1.Val == node2.Val && res1 && res2
	}
	return sameTree(p, q)
}

func main() {
	p := &TreeNode{1, nil, nil}
	p.Left = &TreeNode{2, nil, nil}
	p.Right = &TreeNode{1, nil, nil}
	q := &TreeNode{1, nil, nil}
	q.Left = &TreeNode{1, nil, nil}
	q.Right = &TreeNode{2, nil, nil}
	fmt.Println(isSameTree(p, q))
}
