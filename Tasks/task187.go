package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	m := make(map[string]int)

	for i := 0; i <= len(s)-10; i++ {
		subString := s[i : i+10]
		if m[subString] == 1 {
			result = append(result, subString)
		}
		m[subString]++
	}

	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
