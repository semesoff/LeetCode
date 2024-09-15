package main

import (
	"fmt"
	"strings"
)

type Stack []string

func (s *Stack) add(value string) {
	*s = append(*s, value)
}

func (s *Stack) pop() string {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func makeString(path Stack) string {
	var result string
	for _, j := range path {
		result += fmt.Sprintf("/%s", j)
	}
	return result
}

func simplifyPath(path string) string {
	var current strings.Builder
	stack := Stack{}

	for _, char := range path {
		if char != '/' {
			current.WriteRune(char)
		} else if current.String() == ".." {
			if !stack.IsEmpty() {
				stack.pop()
			}
			current.Reset()
		} else if current.String() == "." {
			current.Reset()
		} else if current.Len() > 0 {
			stack.add(strings.TrimSpace(current.String()))
			current.Reset()
		}
	}
	if current.Len() > 0 {
		if str := current.String(); str != ".." && str != "." {
			stack.add(current.String())
		}
		if current.String() == ".." {
			if !stack.IsEmpty() {
				stack.pop()
			}
		}
	}
	if stack.IsEmpty() {
		return "/"
	}
	return makeString(stack)
}

func main() {
	fmt.Println(simplifyPath("/home/of/foo/../../bar/../../is/./here/."))
}
