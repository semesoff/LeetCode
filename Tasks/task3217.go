package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	hash := make(map[int]int)
	for _, i := range nums {
		hash[i] = 1
	}

	nodes := make([]int, 0)
	node := head
	for node != nil {
		if _, ok := hash[node.Val]; !ok {
			nodes = append(nodes, node.Val)
		}
		node = node.Next
	}

	cur := head
	old := head
	for _, i := range nodes {
		old = cur
		cur.Val = i
		if cur.Next == nil {
			cur.Next = &ListNode{}
		}
		cur = cur.Next
	}
	old.Next = nil
	return head
}

func main() {
	list := &ListNode{1, nil}
	list.Next = &ListNode{2, nil}
	list.Next.Next = &ListNode{3, nil}
	list.Next.Next.Next = &ListNode{4, nil}
	list.Next.Next.Next.Next = &ListNode{5, nil}
	fmt.Println(modifiedList([]int{1, 2, 3}, list).Next.Next)
}
