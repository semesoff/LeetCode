package main

import "fmt"

func lemonadeChange(bills []int) bool {
	countOfFive := 0
	countOfTen := 0
	for _, i := range bills {
		if i == 5 {
			countOfFive++
		} else if i == 10 {
			if countOfFive >= 1 {
				countOfTen++
				countOfFive--
			} else {
				return false
			}
		} else {
			if countOfFive >= 1 && countOfTen >= 1 {
				countOfTen--
				countOfFive--
			} else if countOfFive >= 3 {
				countOfFive -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5, 5, 5, 10, 20}))
}
