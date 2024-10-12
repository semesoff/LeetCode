package main

import "fmt"

func lengthOfLastWord(s string) int {
	res := 0
	isAdd := false
	for i := range s {
		if s[i] != ' ' {
			if isAdd {
				res = 0
				isAdd = false
			}
			res++
		} else {
			isAdd = true
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
