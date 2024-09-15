package main

import (
	"fmt"
	"strconv"
)

type Stack []string

func (s *Stack) add(val string) {
	*s = append(*s, val)
}

func (s *Stack) pop() string {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) getLast() int {
	res, _ := strconv.Atoi((*s)[len(*s)-1])
	return res
}

func (s *Stack) getSecondLast() int {
	res, _ := strconv.Atoi((*s)[len(*s)-2])
	return res
}

func sumResults(s *Stack) int {
	res := 0
	for _, i := range *s {
		num, _ := strconv.Atoi(i)
		res += num
	}
	return res
}

func calPoints(operations []string) int {
	stack := Stack{}
	for _, ch := range operations {
		if ch == "+" {
			stack.add(strconv.Itoa(stack.getLast() + stack.getSecondLast()))
		} else if ch == "D" {
			stack.add(strconv.Itoa(stack.getLast() * 2))
		} else if ch == "C" {
			stack.pop()
		} else {
			stack.add(ch)
		}
	}
	return sumResults(&stack)
}

func main() {
	fmt.Println(calPoints([]string{"5", "-2", "4", " C", "D", "9", "+", "+"}))
}
