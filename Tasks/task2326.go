package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func changeCord(i, j, vector *int, m, n int, matrix [][]int) {
	switch *vector {
	case 1:
		if *i-1 < 0 || matrix[*i-1][*j] != -1 {
			*j++
			*vector = 2
		} else {
			*i--
		}
	case 2:
		if *j+1 == n || matrix[*i][*j+1] != -1 {
			*i++
			*vector = 3
		} else {
			*j++
		}
	case 3:
		if *i+1 == m || matrix[*i+1][*j] != -1 {
			*j--
			*vector = 4
		} else {
			*i++
		}
	case 4:
		if *j-1 < 0 || matrix[*i][*j-1] != -1 {
			*i--
			*vector = 1
		} else {
			*j--
		}
	}
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	node := head
	vector := 2
	i, j := 0, 0
	for node != nil {
		matrix[i][j] = node.Val
		changeCord(&i, &j, &vector, m, n, matrix)
		node = node.Next
	}
	return matrix
}

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 0}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	fmt.Println(spiralMatrix(3, 5, head))
}
