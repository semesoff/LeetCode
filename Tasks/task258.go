package main

import "fmt"

func sumDigitOfNumber(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func addDigits(num int) int {
	for num/10 != 0 {
		num = sumDigitOfNumber(num)
	}
	return num
}

func main() {
	fmt.Println(addDigits(38))
}
