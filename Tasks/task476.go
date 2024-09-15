package main

import (
	"fmt"
	"strconv"
)

func swapNumbers(number string) string {
	bytes := []byte(number)
	for i := range bytes {
		if bytes[i] == '0' {
			bytes[i] = '1'
		} else {
			bytes[i] = '0'
		}
	}
	return string(bytes)
}

func findComplement(num int) int {
	formatInt := strconv.FormatInt(int64(num), 2)
	res, _ := strconv.ParseInt(swapNumbers(formatInt), 2, 64)
	return int(res)
}

func main() {
	fmt.Println(findComplement(5))
}
