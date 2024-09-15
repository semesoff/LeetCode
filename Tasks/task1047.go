package main

import (
	"fmt"
)

type Stack []rune

func (s *Stack) add(value rune) {
	*s = append(*s, value)
}

func (s *Stack) pop() rune {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) getLast() rune {
	return (*s)[len(*s)-1]
}

func removeDuplicates(s string) string {
	stack := Stack{}
	for _, ch := range s {
		if !stack.isEmpty() && stack.getLast() == ch {
			stack.pop()
		} else {
			stack.add(ch)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("azxxzy"))
}
