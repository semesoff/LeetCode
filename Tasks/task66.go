package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1]++

	if digits[n-1] > 9 {
		remainder := 0
		for i := n - 1; i >= 0; i-- {
			digits[i] += remainder
			remainder = 0
			if digits[i] > 9 {
				remainder = digits[i] / 10
				digits[i] %= 10
			}
			if i == 0 && remainder > 0 {
				digits = append([]int{remainder}, digits...)
			}
		}
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 9})) // [1, 2, 4]
}
