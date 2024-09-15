package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := make([]int, 0)
	var recursion func(node *Node)
	recursion = func(node *Node) {
		if node == nil {
			return
		}
		for _, i := range node.Children {
			recursion(i)
		}
		res = append(res, node.Val)
	}
	recursion(root)
	return res
}

func main() {
	tree := &Node{}
}
