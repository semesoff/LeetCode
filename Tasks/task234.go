package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}

	cur = head
	for cur != nil && cur.Val == stack[len(stack)-1] {
		cur = cur.Next
		stack = stack[:len(stack)-1]
	}
	return cur == nil
}

func main() {
	head := &ListNode{1, nil}
	head.Next = &ListNode{1, nil}
	head.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next = &ListNode{1, nil}
	fmt.Println(isPalindrome(head))
}
