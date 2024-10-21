package main

import "fmt"

func getNumber(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	visited := make(map[int]bool)

	for _, ok := visited[n]; !ok; _, ok = visited[n] {
		visited[n] = true
		n = getNumber(n)
		if n == 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isHappy(1))
}
