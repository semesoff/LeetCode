package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	nodes := make([]int, 0)
	size := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		nodes = append(nodes, node.Val)
		size++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	sort.Ints(nodes)
	left, right := 0, size-1
	for left < right {
		if nodes[left]+nodes[right] == k {
			return true
		} else if nodes[left]+nodes[right] < k {
			left++
		} else {
			right--
		}

	}
	return false
}

func main() {
	fmt.Println(findTarget())
}
