package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mapToSlice(x map[int]int) []int {
	res := make([]int, 0)
	for i := range x {
		res = append(res, i)
	}
	sort.Ints(res)
	return res
}

func deleteDuplicates(head *ListNode) *ListNode {
	hash := make(map[int]int)
	unique := make(map[int]int)
	node := head
	for node != nil {
		if _, ok := hash[node.Val]; ok {
			delete(unique, node.Val)
		} else {
			hash[node.Val] = 1
			unique[node.Val] = 1
		}
		node = node.Next
	}

	cur := head
	old := &ListNode{}
	if len(unique) == 0 {
		return nil
	}
	for _, i := range mapToSlice(unique) {
		cur.Val = i
		if cur.Next == nil {
			cur.Next = &ListNode{}
		}
		old = cur
		cur = cur.Next
	}
	old.Next = nil
	return head
}

func main() {
	list := &ListNode{}
	fmt.Println(deleteDuplicates(list))
}
