package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mapToValues(x map[int]*TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	for _, v := range x {
		res = append(res, v)
	}
	return res
}

func mapToSlice(x map[int]*TreeNode) []int {
	res := make([]int, 0)
	for k, v := range x {
		if v == nil {
			res = append(res, k)
		}
	}
	return res
}

func find(x int, slice []int) bool {
	for _, v := range slice {
		if x == v {
			return true
		}
	}
	return false
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	res := map[int]*TreeNode{root.Val: root}
	var dfs func(*TreeNode, *TreeNode, bool)

	dfs = func(node *TreeNode, parent *TreeNode, isLeft bool) {
		if node == nil {
			return
		}

		dfs(node.Left, node, true)
		dfs(node.Right, node, false)

		if find(node.Val, toDelete) {
			if find(node.Val, mapToSlice(res)) {
				delete(res, node.Val)
			}

			if parent != nil {
				if isLeft {
					parent.Left = nil
				} else {
					parent.Right = nil
				}
			}

			if node.Left != nil {
				res[node.Left.Val] = node.Left
			}
			if node.Right != nil {
				res[node.Right.Val] = node.Right
			}
		}
	}
	dfs(root, nil, false)
	return mapToValues(res)
}

func main() {
	tree := &TreeNode{1, nil, nil}
	tree.Left = &TreeNode{2, nil, nil}
	tree.Right = &TreeNode{3, nil, nil}
	tree.Left.Left = &TreeNode{4, nil, nil}
	tree.Left.Right = &TreeNode{5, nil, nil}
	tree.Right.Left = &TreeNode{6, nil, nil}
	tree.Right.Right = &TreeNode{7, nil, nil}
	fmt.Println(delNodes(tree, []int{3, 5}))
}
