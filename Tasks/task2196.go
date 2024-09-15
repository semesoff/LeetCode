package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type IntSlice []int

func (slice IntSlice) isExist(value int) bool {
	for _, i := range slice {
		if i == value {
			return true
		}
	}
	return false
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := map[int]*TreeNode{}
	children := IntSlice{}

	for _, data := range descriptions {
		parent := data[0]
		child := data[1]
		isLeft := data[2]

		if _, err := nodes[parent]; !err {
			nodes[parent] = &TreeNode{Val: parent, Left: nil, Right: nil}
		}
		if _, err := nodes[child]; !err {
			nodes[child] = &TreeNode{Val: child, Left: nil, Right: nil}
		}
		children = append(children, child)

		if isLeft == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
	}

	var root *TreeNode

	for _, node := range descriptions {
		if !children.isExist(node[0]) {
			root = nodes[node[0]]
			break
		}
	}
	return root
}

func main() {
	fmt.Println(createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}))
}
