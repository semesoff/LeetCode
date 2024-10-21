package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords(" the   sky is blue")) // "blue is sky the"
}
