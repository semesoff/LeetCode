package main

import (
	"fmt"
	"strconv"
)

func isGood(num string) bool {
	byteStr := []rune(num)
	res1 := 0
	res2 := 0
	n := len(num)
	for i := 0; i < n; i++ {
		if i < n/2 {
			res1 += int(byteStr[i] - '0')
		} else {
			res2 += int(byteStr[i] - '0')
		}
	}

	return res1 == res2
}

func countSymmetricIntegers(low int, high int) int {
	total := 0
	for i := low; i <= high; i++ {
		strNum := strconv.Itoa(i)
		if len(strNum)%2 == 0 && isGood(strNum) {
			total++
		}
	}
	return total
}

func main() {
	fmt.Println(countSymmetricIntegers(1, 100))
}
