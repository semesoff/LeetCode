package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode, strNum string)
	dfs = func(node *TreeNode, strNum string) {
		if node == nil {
			return
		}
		strNum += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			num, _ := strconv.ParseInt(strNum, 2, 64)
			res += int(num)
		}
		dfs(node.Left, strNum)
		dfs(node.Right, strNum)
	}
	dfs(root, "")
	return res
}

func main() {
	tree := &TreeNode{1, nil, nil}
	tree.Left = &TreeNode{0, nil, nil}
	tree.Left.Left = &TreeNode{0, nil, nil}
	tree.Left.Right = &TreeNode{1, nil, nil}
	tree.Right = &TreeNode{1, nil, nil}
	tree.Right.Right = &TreeNode{1, nil, nil}
	tree.Right.Left = &TreeNode{0, nil, nil}
	fmt.Println(sumRootToLeaf(tree))
}
