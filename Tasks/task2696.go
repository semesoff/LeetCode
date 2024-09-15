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

func minLength(s string) int {
	stack := Stack{}
	total := 0
	for _, ch := range s {
		if !stack.isEmpty() && stack.getLast() == 'A' && ch == 'B' {
			stack.pop()
			total--
		} else if !stack.isEmpty() && stack.getLast() == 'C' && ch == 'D' {
			stack.pop()
			total--
		} else {
			stack.add(ch)
			total++
		}
	}
	return total
}

func main() {
	fmt.Println(minLength("ABFCACDB"))
}
