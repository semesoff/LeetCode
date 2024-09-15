package main

import "fmt"

func distMoney(money int, children int) int {
	total := 0
	if money < children {
		return -1
	}
	if money > children*8 {
		return children - 1
	}

	for money > 0 && money-8 >= children-1 {
		total++
		money -= 8
		children--
	}
	if money == 4 && children == 1 {
		total--
	}
	return total
}

func main() {
	fmt.Println(distMoney(1, 2))
}
