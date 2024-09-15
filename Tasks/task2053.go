package main

import (
	"fmt"
)

func kthDistinct(arr []string, k int) string {
	hash := map[string]int{}
	for _, i := range arr {
		hash[i]++
	}
	ku := 1
	for _, val := range arr {
		if hash[val] == 1 && ku == k {
			return val
		} else if hash[val] == 1 {
			ku++
		}
	}
	return ""
}

func main() {
	fmt.Println(kthDistinct([]string{"d", "b", "c", "b", "c", "a"}, 2))
}
