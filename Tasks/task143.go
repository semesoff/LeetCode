package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	var previous *ListNode
	last := head

	for last.Next != nil {
		previous = last
		last = last.Next
	}
	previous.Next = nil

	tmp := head.Next
	head.Next = last
	last.Next = tmp
	reorderList(tmp)
}

func main() {
	tree := &ListNode{1, nil}
	tree.Next = &ListNode{2, nil}
	tree.Next.Next = &ListNode{3, nil}
	tree.Next.Next.Next = &ListNode{4, nil}
	reorderList(tree)
}
