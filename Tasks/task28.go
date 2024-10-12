package main

import "fmt"

func strStr(haystack string, needle string) int {
	n := len(haystack)
	k := len(needle)
	res := -1
	for i := 0; i < n; i++ {
		if i+k > n {
			break
		}
		if haystack[i:i+k] == needle {
			res = i
			break
		}
	}
	return res
}

func main() {
	fmt.Println(strStr("a", "a"))
}
