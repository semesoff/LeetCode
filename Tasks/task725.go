package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listSize(list *ListNode) int {
	n := 0
	cur := list
	for cur != nil {
		n++
		cur = cur.Next
	}
	return n
}

//func splitListToParts(head *ListNode, k int) []*ListNode {
//	sub := make([]*ListNode, k)
//	n := listSize(head)
//
//	size := n / k
//	extra := n % k
//
//	cur := head
//	for i := range sub {
//		arr := &ListNode{}
//
//		localSize := size
//		if extra > 0 {
//			localSize++
//			extra--
//		}
//		curSize := localSize
//		arrCop := arr
//		old := arr
//		for localSize != 0 {
//			old = arrCop
//			if cur == nil {
//				break
//			}
//			arrCop.Val = cur.Val
//			if arrCop.Next == nil {
//				arrCop.Next = &ListNode{}
//			}
//			arrCop = arrCop.Next
//			localSize--
//			cur = cur.Next
//		}
//		old.Next = nil
//		if curSize != 0 {
//			sub[i] = arr
//		}
//	}
//	return sub
//}

func splitListToParts(head *ListNode, k int) []*ListNode {
	sub := make([]*ListNode, k)
	n := listSize(head)

	size := n / k
	extra := n % k

	cur := head
	for i := range sub {
		arr := &ListNode{}
		localSize := size

		if extra > 0 {
			localSize++
			extra--
		}

		if localSize == 0 {
			continue
		}

		arrCop := arr
		old := arr
		for localSize != 0 {
			old = arrCop
			if cur == nil {
				break
			}
			arrCop.Val = cur.Val
			if arrCop.Next == nil {
				arrCop.Next = &ListNode{}
			}
			arrCop = arrCop.Next
			localSize--
			cur = cur.Next
		}
		old.Next = nil
		sub[i] = arr
	}
	return sub
}

func main() {
	list := &ListNode{1, nil}
	list.Next = &ListNode{2, nil}
	list.Next.Next = &ListNode{3, nil}
	fmt.Println(splitListToParts(list, 5))
}
