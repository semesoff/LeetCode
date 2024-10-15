package main

import "fmt"

func minimumSteps(s string) int64 {
	steps := 0
	n := len(s)
	lastBlack := -1

	for i := n - 1; i > -1; i-- {
		if s[i] == '1' && lastBlack == -1 {
			lastBlack = n - 1
			steps += lastBlack - i
		} else if s[i] == '1' {
			steps += lastBlack - i - 1
			lastBlack--
		}
	}
	return int64(steps)
}

func main() {
	fmt.Println(minimumSteps("11000111"))
}
