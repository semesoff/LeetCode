package main

import (
	"fmt"
)

//func largestRectangleArea(heights []int) int {
//	maxArea := -math.MaxInt32
//	n := len(heights)
//	var dfs func(count int, minNumber, index int)
//	dfs = func(count int, minNumber, index int) {
//		if index == n {
//			return
//		}
//		dfs(0, math.MaxInt32, index+1)
//		count++
//		if heights[index] < minNumber {
//			minNumber = heights[index]
//		}
//		dfs(count, minNumber, index+1)
//		if res := minNumber * count; res > maxArea {
//			maxArea = res
//		}
//	}
//	dfs(0, math.MaxInt32, 0)
//	return maxArea
//}

type Stack []int

func (s *Stack) add(value int) {
	*s = append(*s, value)
}

func (s *Stack) pop() int {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) getLast() int {
	return (*s)[len(*s)-1]
}

// Монотонный стек
func largestRectangleArea(heights []int) int {
	stack := Stack{}
	maxArea := 0
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		for !stack.isEmpty() && heights[i] < heights[stack.getLast()] {
			h := heights[stack.pop()]
			width := i
			if !stack.isEmpty() {
				width = i - stack.getLast() - 1
			}
			maxArea = max(maxArea, h*width)
		}
		stack.add(i)
	}
	return maxArea
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}
