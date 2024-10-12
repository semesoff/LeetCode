package main

import "fmt"

func isAlphanumeric(c byte) bool {
	return (c >= '0' && c <= '9') || (toLower(c) >= 'a' && toLower(c) <= 'z')
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
