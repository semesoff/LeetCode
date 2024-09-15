package main

import "fmt"

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

func isGood(open, close rune) bool {
	switch open {
	case '(':
		if close == ')' {
			return true
		}
	case '[':
		if close == ']' {
			return true
		}
	case '{':
		if close == '}' {
			return true
		}
	}
	return false
}

func isValid(s string) bool {
	stack := Stack{}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack.add(ch)
		} else {
			if !stack.isEmpty() && isGood(stack.getLast(), ch) {
				stack.pop()
			} else {
				return false
			}
		}
	}

	return stack.isEmpty()
}

func main() {
	fmt.Println(isValid("{(())}"))
}
