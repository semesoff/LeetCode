package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var findDepth func(node *TreeNode, depth int) int
	findDepth = func(node *TreeNode, depth int) int {
		if node != nil {
			res1 := findDepth(node.Left, depth+1)
			res2 := findDepth(node.Right, depth+1)
			if res1 >= res2 {
				depth = res1
			} else {
				depth = res2
			}
		}
		return depth
	}
	return findDepth(root, 0)
}
