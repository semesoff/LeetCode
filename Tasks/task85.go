package main

import "fmt"

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

func sumSlices(slice1 *[]int, slice2 []byte) {
	for i := range *slice1 {
		if slice2[i] == '0' {
			(*slice1)[i] = 0
		} else {
			(*slice1)[i] += int(slice2[i] - '0')
		}
	}
}

func calculateMaxArea(slice []int, maxArea *int) {
	stack := Stack{}
	for i := 0; i <= len(slice); i++ {
		h := 0
		if i < len(slice) {
			h = slice[i]
		}
		for !stack.isEmpty() && slice[stack.getLast()] >= h {
			height := slice[stack.pop()]
			width := i
			if !stack.isEmpty() {
				width = i - stack.getLast() - 1
			}
			*maxArea = max(*maxArea, height*width)
		}
		stack.add(i)
	}
}

func maximalRectangle(matrix [][]byte) int {
	c := len(matrix[0])
	histogram := make([]int, c)
	maxArea := 0

	for _, h := range matrix {
		sumSlices(&histogram, h)
		calculateMaxArea(histogram, &maxArea)
	}
	return maxArea
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
