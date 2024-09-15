package main

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	res := 0
	for _, i := range s {
		number := int(rune(i)-'a') + 1
		res += number/10 + number%10
	}

	for i := 0; i < k-1; i++ {
		str := strconv.Itoa(res)
		res = 0
		for _, j := range str {
			number, _ := strconv.Atoi(string(j))
			res += number
		}
	}
	return res
}

func main() {
	fmt.Println(getLucky("leetcode", 2))
}
