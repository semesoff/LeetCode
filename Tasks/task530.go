package main

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func getMinimumDifference(root *TreeNode) int {
//	nodes := make([]int, 0)
//	var dfs func(node *TreeNode)
//	dfs = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		nodes = append(nodes, node.Val)
//		dfs(node.Left)
//		dfs(node.Right)
//	}
//	dfs(root)
//	sort.Ints(nodes)
//	n := len(nodes)
//	minDifference := math.MaxInt32
//	for i := 0; i < n-1; i++ {
//		minDifference = min(minDifference, nodes[i+1]-nodes[i])
//	}
//	return minDifference
//}

func getMinimumDifference(root *TreeNode) int {
	x := math.MaxInt32
	minDiff := math.MaxInt32
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		minDiff = min(minDiff, int(math.Abs(float64(node.Val-x))))
		x = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return minDiff
}
