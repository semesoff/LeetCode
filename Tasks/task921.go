package main

import (
	"fmt"
)

//type Stack []rune
//
//func (s *Stack) Push(r rune) {
//	*s = append(*s, r)
//}
//
//func (s *Stack) Pop() rune {
//	item := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return item
//}
//
//func (s *Stack) Top() rune {
//	return (*s)[len(*s)-1]
//}
//
//func (s *Stack) isEmpty() bool {
//	return len(*s) == 0
//}
//
//func minAddToMakeValid(s string) int {
//	stack := Stack{}
//	for _, ch := range s {
//		if !stack.isEmpty() && stack.Top() == '(' && ch == ')' {
//			stack.Pop()
//		} else {
//			stack.Push(ch)
//		}
//	}
//	return len(stack)
//}

func minAddToMakeValid(s string) int {
	opened := 0
	closed := 0
	for _, ch := range s {
		if opened > 0 && ch == ')' {
			opened--
		} else {
			if ch == '(' {
				opened++
			} else {
				closed++
			}
		}
	}
	return opened + closed
}

func main() {
	fmt.Println(minAddToMakeValid("()))"))
}
