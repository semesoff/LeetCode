package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyLetters(s string) string {
	res := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !unicode.IsLetter(rune(s[i])) {
			i++
		}
		for i < j && !unicode.IsLetter(rune(s[j])) {
			j--
		}
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}

func main() {
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
