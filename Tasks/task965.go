package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	var dfs func(node *TreeNode, value int, isUniValued *bool)
	dfs = func(node *TreeNode, value int, isUniValued *bool) {
		if node == nil {
			return
		}
		if node.Val != value {
			*isUniValued = false
			return
		}
		value = node.Val
		dfs(node.Left, value, isUniValued)
		dfs(node.Right, value, isUniValued)
	}
	isUniValued := true
	dfs(root, root.Val, &isUniValued)
	return isUniValued
}

func main() {

}
