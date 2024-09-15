package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createPointer(list, additional *ListNode) *ListNode {
	node := list
	for node != nil {
		if node.Next == nil {
			node.Next = additional
			break
		}
		node = node.Next
	}
	return list
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	next := &ListNode{}
	previous := list1
	node := list1
	k := 0
	for node != nil {
		if k < a {
			previous = node
		} else if k-1 == b {
			next = node
		}
		node = node.Next
		k++
	}
	previous.Next = createPointer(list2, next)
	return list1
}

func main() {
	// Создание list1 = [0,1,2,3,4,5,6]
	list1 := &ListNode{Val: 0}
	list1.Next = &ListNode{Val: 1}
	list1.Next.Next = &ListNode{Val: 2}
	list1.Next.Next.Next = &ListNode{Val: 3}
	list1.Next.Next.Next.Next = &ListNode{Val: 4}
	list1.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	list1.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	// Создание list2 = [1000000,1000001,1000002,1000003,1000004]
	list2 := &ListNode{Val: 1000000}
	list2.Next = &ListNode{Val: 1000001}
	list2.Next.Next = &ListNode{Val: 1000002}
	list2.Next.Next.Next = &ListNode{Val: 1000003}
	list2.Next.Next.Next.Next = &ListNode{Val: 1000004}
	fmt.Println(mergeInBetween(list1, 2, 5, list2).Next.Next.Next.Next.Next.Next.Next)
}
