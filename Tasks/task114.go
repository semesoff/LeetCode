package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	nums := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nums = append(nums, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	node := root
	node.Left = nil
	node.Right = &TreeNode{}
	for i := 1; i < len(nums); i++ {
		node = node.Right
		node.Val = nums[i]
		node.Left = nil
		node.Right = &TreeNode{}
	}
	node.Right = nil
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}
	flatten(root)
}
