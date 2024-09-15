package main

func preorder(root *Node) []int {
	res := make([]int, 0)
	var recursion func(node *Node)
	recursion = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, i := range node.Children {
			recursion(i)
		}
	}
	recursion(root)
	return res
}

func main() {

}
