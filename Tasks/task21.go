package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	node1 := list1
	node2 := list2
	res := &ListNode{}
	pointer := res
	for node1 != nil || node2 != nil {
		if node1 != nil && node2 == nil || node1 != nil && node1.Val < node2.Val {
			pointer.Val = node1.Val
			node1 = node1.Next
		} else {
			pointer.Val = node2.Val
			node2 = node2.Next
		}
		if pointer.Next == nil && node1 != nil || node2 != nil {
			pointer.Next = &ListNode{}
		}
		pointer = pointer.Next
	}
	return res
}

func main() {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	fmt.Println(mergeTwoLists(list1, list2))
}
