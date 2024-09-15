package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, val, depth, currDepth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root, Right: nil}
	}
	if root == nil {
		return root
	}
	if currDepth == depth-1 {
		leftNode := root.Left
		rightNode := root.Right
		root.Left = &TreeNode{Val: val, Left: leftNode, Right: nil}
		root.Right = &TreeNode{Val: val, Left: nil, Right: rightNode}
		return root
	}
	helper(root.Left, val, depth, currDepth+1)
	helper(root.Right, val, depth, currDepth+1)
	return root
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	return helper(root, val, depth, 1)
}

func main() {

}
