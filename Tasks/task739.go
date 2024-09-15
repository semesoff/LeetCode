package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Pop() int {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) Add(value int) {
	*s = append(*s, value)
}

func dailyTemperatures(temperatures []int) []int {
	stack := Stack{}
	res := make([]int, len(temperatures))
	for i := range temperatures {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			index := stack.Pop()
			res[index] = i - index
		}
		stack.Add(i)
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
