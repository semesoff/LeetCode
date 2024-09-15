package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack []*ListNode

func (s *Stack) pop() *ListNode {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) add(val *ListNode) {
	*s = append(*s, val)
}

func (s *Stack) getLast() *ListNode {
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func removeNodes(head *ListNode) *ListNode {
	stack := Stack{}
	cur := head
	for cur != nil {
		for !stack.isEmpty() && stack.getLast().Val < cur.Val {
			stack.pop()
			if !stack.isEmpty() {
				stack.getLast().Next = cur
			}
		}
		stack.add(cur)
		cur = cur.Next
	}
	return stack[0]
}

func main() {
	tree := &ListNode{5, nil}
	tree.Next = &ListNode{2, nil}
	tree.Next.Next = &ListNode{13, nil}
	tree.Next.Next.Next = &ListNode{3, nil}
	tree.Next.Next.Next.Next = &ListNode{8, nil}
	fmt.Println(removeNodes(tree))
}
