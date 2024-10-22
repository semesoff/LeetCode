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

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return node
}

func (s *Stack) PopLeft() *TreeNode {
	node := (*s)[0]
	*s = (*s)[1:]
	return node
}

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	stack := Stack{root}
	sums := make([]int, 0)
	sums = append(sums, root.Val)

	for !stack.isEmpty() {
		n := len(stack)
		for i := 0; i < n; i++ {
			node := stack.PopLeft()
			if node.Left != nil {
				stack.Push(node.Left)
			}
			if node.Right != nil {
				stack.Push(node.Right)
			}
		}
		curSum := 0
		for i := 0; i < len(stack); i++ {
			curSum += stack[i].Val
		}
		if curSum != 0 {
			sums = append(sums, curSum)
		}
	}

	if len(sums) < k {
		return -1
	}
	sort.Ints(sums)
	return int64(sums[len(sums)-k])
}

func main() {
	tree := &TreeNode{1, nil, nil}
	tree.Left = &TreeNode{2, nil, nil}
	tree.Left.Left = &TreeNode{3, nil, nil}

	fmt.Println(kthLargestLevelSum(tree, 1))
}
