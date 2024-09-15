package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Element struct {
	val   int
	index int
}

type Stack []Element

func (s *Stack) add(value Element) {
	*s = append(*s, value)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func getSize(head *ListNode) int {
	size := 0
	for head != nil {
		head = head.Next
		size++
	}
	return size
}

func (s *Stack) pop() Element {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) getLast() Element {
	return (*s)[len(*s)-1]
}

func nextLargerNodes(head *ListNode) []int {
	res := make([]int, getSize(head))
	stack := Stack{}
	node := head
	index := 0
	for node != nil {
		for !stack.isEmpty() && stack.getLast().val < node.Val {
			top := stack.pop()
			res[top.index] = node.Val
		}
		stack.add(Element{val: node.Val, index: index})
		index++
		node = node.Next
	}
	return res
}

// 1,7,5,1,9,2,5,1
func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{7, nil}
	head.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next = &ListNode{1, nil}
	head.Next.Next.Next.Next = &ListNode{9, nil}
	head.Next.Next.Next.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{1, nil}
	fmt.Println(nextLargerNodes(head))
}
