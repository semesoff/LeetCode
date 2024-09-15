package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func findBottomLeftValue(root *TreeNode) int {
//	maxLength, res := 0, 0
//	var dfs func(node *TreeNode, length int)
//	dfs = func(node *TreeNode, length int) {
//		if node == nil {
//			return
//		}
//		if node.Left == node.Right {
//			if length >= maxLength {
//				maxLength = length
//				res = node.Val
//			}
//		}
//		dfs(node.Left, length+1)
//		dfs(node.Right, length+1)
//	}
//	dfs(root, 0)
//	return res
//}

func findBottomLeftValue(root *TreeNode) int {
	var queue []*TreeNode
	queue = append(queue, root)
	var res int

	for len(queue) > 0 {
		node := queue[0]
		queue := queue[1:]

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		res = node.Val
	}

	return res
}

func main() {

}
