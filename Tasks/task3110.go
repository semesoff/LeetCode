package main

import (
	"fmt"
	"math"
)

func scoreOfString(s string) int {
	n := len(s)
	result := 0

	for i := 1; i < n; i++ {
		result += int(math.Abs(float64(int(s[i-1]) - int(s[i]))))
	}

	return result
}

func main() {
	result := scoreOfString("hello")
	fmt.Println(result)
}
