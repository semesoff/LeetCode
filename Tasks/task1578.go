package main

import "fmt"

type Rope struct {
	Color rune
	Time  int
}

// 1. Solution with Stack
//type Stack []*Rope
//
//func (s *Stack) Push(r *Rope) {
//	*s = append(*s, r)
//}
//
//func (s *Stack) Pop() Rope {
//	r := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return *r
//}
//
//func (s *Stack) Top() *Rope {
//	return (*s)[len(*s)-1]
//}
//
//func (s *Stack) IsEmpty() bool {
//	return len(*s) == 0
//}
//
//func minCost(color string, neededTime []int) int {
//	n := len(neededTime)
//	res := 0
//
//	stack := Stack{}
//	curMax := 0
//	curSum := 0
//	str := []byte(color)
//	for i, ch := range str {
//		if !stack.IsEmpty() && stack.Top().Color == rune(ch) {
//			curSum += stack.Top().Time
//			curMax = max(curMax, stack.Top().Time)
//			stack.Pop()
//			if i == n-1 || (i < n-1 && str[i+1] != ch) {
//				curSum += neededTime[i]
//				curMax = max(curMax, neededTime[i])
//				res += curSum - curMax
//				curMax = 0
//				curSum = 0
//			}
//		}
//		stack.Push(&Rope{Color: rune(ch), Time: neededTime[i]})
//	}
//	return res
//}

// 2. Solution without Stack (with O(1) space allocation)
func minCost(color string, neededTime []int) int {
	n := len(neededTime)
	res := 0
	curMax := 0
	curSum := 0

	previousRope := Rope{Color: 'a', Time: 0}
	str := []byte(color)
	for i, ch := range str {
		if previousRope.Time != 0 && previousRope.Color == rune(ch) {
			curSum += previousRope.Time
			curMax = max(curMax, previousRope.Time)
			previousRope.Time = 0
			if i == n-1 || (i < n-1 && str[i+1] != ch) {
				curSum += neededTime[i]
				curMax = max(curMax, neededTime[i])
				res += curSum - curMax
				curMax = 0
				curSum = 0
			}
		}
		previousRope = Rope{Color: rune(ch), Time: neededTime[i]}
	}
	return res
}

func main() {
	fmt.Println(minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}))
}

// 8 8 10
