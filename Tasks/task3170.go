package main

import (
	"fmt"
	"sort"
	"strings"
)

type StackList []rune
type Stack map[rune][]int

func (s *StackList) add(value rune) {
	*s = append(*s, value)
}

func (s *StackList) pop(index int) rune {
	(*s)[index] = '#'
	return ' '
}

func (s *Stack) add(value rune, index int) {
	if _, ok := (*s)[value]; ok {
		(*s)[value] = append((*s)[value], index)
		sort.Ints((*s)[value])
	} else {
		(*s)[value] = []int{index}
	}
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) getChar() rune {
	char := 'z'
	for i, _ := range *s {
		if i < char {
			char = i
		}
	}
	return char
}

func (s *Stack) pop(key rune) int {
	item := (*s)[key][len((*s)[key])-1]
	(*s)[key] = (*s)[key][:len((*s)[key])-1]
	if len((*s)[key]) == 0 {
		delete(*s, key)
	}
	return item
}

func deleteChar(s *string) string {
	stack := Stack{}
	stackList := StackList{}
	step := 0
	for idx, i := range *s {
		if i == '*' {
			step++
			if !stack.isEmpty() {
				stackList.pop(stack.pop(stack.getChar()))
			}
		} else {
			stackList.add(i)
			stack.add(i, idx-step)
		}
	}

	result := ""
	for _, i := range stackList {
		if i != '#' {
			result = strings.Join([]string{result, string(i)}, "")
		}

	}
	return result
}

func clearStars(s string) string {
	return deleteChar(&s)
}

func main() {
	fmt.Println(clearStars("aaba*"))
}
