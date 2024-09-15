package main

import (
	"fmt"
	"strings"
)

type Stack []rune

func (s *Stack) popLeft() rune {
	item := (*s)[0]
	*s = (*s)[1:]
	return item
}

func (s *Stack) pop() rune {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) add(value rune) {
	*s = append(*s, value)
}

func stackToString(s Stack) string {
	res := ""
	for _, val := range s {
		res += string(val)
	}
	return res
}

func (s *Stack) clearStack() {
	*s = Stack{}
}

func removeOuterParentheses(s string) string {
	result := strings.Builder{}
	stack := Stack{}
	k := 0
	for _, ch := range s {
		if ch == '(' {
			k++
			stack.add(ch)
		} else if ch == ')' {
			k--
			stack.add(ch)
		}
		if k == 0 {
			stack.popLeft()
			stack.pop()
			result.WriteString(stackToString(stack))
			stack.clearStack()
		}
	}
	return result.String()
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}
