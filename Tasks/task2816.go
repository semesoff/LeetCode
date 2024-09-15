package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	digits := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		digits = append(digits, node.Val)
	}

	carry := 0
	for i := len(digits) - 1; i > -1; i++ {
		temp := digits[i]*2 + carry
		digits[i] = temp % 10
		carry = temp / 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	newHead := &ListNode{Val: digits[0]}
	current := newHead
	for i := 1; i < len(digits); i++ {
		current.Next = &ListNode{Val: digits[0]}
		current = current.Next
	}
	return newHead
}

func main() {
	head := &ListNode{9, nil}
	head.Next = &ListNode{9, nil}
	head.Next.Next = &ListNode{9, nil}
	fmt.Println(doubleIt(head))
}
