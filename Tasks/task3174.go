package main

import (
	"fmt"
	"unicode"
)

type Stack []rune

func (s *Stack) pop() rune {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) add(val rune) {
	*s = append(*s, val)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func clearDigits(s string) string {
	stack := Stack{}
	for _, ch := range s {
		if unicode.IsDigit(ch) && !stack.isEmpty() {
			stack.pop()
		} else {
			stack.add(ch)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(clearDigits("cb34"))
}
