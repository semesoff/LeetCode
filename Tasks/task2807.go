package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getGreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	for node := head; node != nil; node = node.Next.Next {
		if node.Next == nil {
			break
		}
		first := node.Val
		second := node.Next.Val
		node.Next = &ListNode{Val: getGreatestCommonDivisor(first, second), Next: node.Next}
	}
	return head
}

func main() {

}
