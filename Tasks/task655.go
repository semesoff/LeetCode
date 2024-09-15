package main

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func heightTree(root *TreeNode) int {
	var height func(node *TreeNode, cur int) int
	height = func(node *TreeNode, cur int) int {
		if node == nil {
			return cur - 1
		}
		return max(height(node.Left, cur+1), height(node.Right, cur+1))
	}
	return height(root, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printTree(root *TreeNode) [][]string {
	height := heightTree(root)
	n := int(math.Pow(2, float64(height)) - 1)
	res := make([][]string, height)
	for i := range res {
		res[i] = make([]string, n)
	}
	var dfs func(node *TreeNode, r, c int)
	dfs = func(node *TreeNode, r, c int) {
		if node == nil {
			return
		}
		res[r][c] = strconv.Itoa(node.Val)
		dfs(node.Left, r+1, c-int(math.Pow(2, float64(height-r-1-1))))
		dfs(node.Right, r+1, c+int(math.Pow(2, float64(height-r-1-1))))
	}
	dfs(root, 0, (n-1)/2)
	return res
}

func main() {
	tree := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(printTree(tree))
}
