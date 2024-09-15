package main

import (
	"fmt"
)

func pop(slice *[]string) string {
	if len(*slice) == 0 {
		return ""
	}
	last := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]
	return last
}

func reverse(s *string) {
	bytes := []byte(*s)
	l, r := 0, len(bytes)-1

	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
	*s = string(bytes)
}

func reverseParentheses(s string) string {
	var queue []string
	var current string

	for _, char := range s {
		if char == '(' {
			queue = append(queue, current)
			current = ""
		} else if char == ')' {
			reverse(&current)
			previous := pop(&queue)
			current = previous + current
		} else {
			current += string(char)
		}
	}
	return current
}

func main() {
	result := reverseParentheses("(ed(et(oc))el)")
	fmt.Println(result)
}
