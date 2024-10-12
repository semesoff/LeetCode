package main

import "fmt"

func minSwaps(s string) int {
	balance := 0
	for _, ch := range s {
		if ch == '[' {
			balance++
		} else {
			if balance > 0 {
				balance--
			}
		}
	}
	return (balance + 1) / 2
}

func main() {
	fmt.Println(minSwaps("]]][[[["))
}
