package main

import (
	"fmt"
	"strings"
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

func (s *Stack) last() rune {
	return (*s)[len(*s)-1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func removeKdigits(num string, k int) string {
	stack := Stack{}
	for _, ch := range num {
		for !stack.isEmpty() && stack.last() > ch && k > 0 {
			stack.pop()
			k--
		}
		stack.add(ch)
	}
	for k > 0 {
		stack.pop()
		k--
	}
	if res := strings.TrimLeft(string(stack), "0"); len(res) > 0 {
		return res
	}
	return "0"
}

func main() {
	fmt.Println(removeKdigits("10", 2))
}
