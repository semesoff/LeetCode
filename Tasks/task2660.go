package main

import "fmt"

func isWinner(player1 []int, player2 []int) int {
	score1 := 0
	score2 := 0
	for i, j := range player1 {
		score1 += j
		score2 += player2[i]
		if i-1 >= 0 && player1[i-1] == 10 || i-2 >= 0 && player1[i-2] == 10 {
			score1 += j
		}
		if i-1 >= 0 && player2[i-1] == 10 || i-2 >= 0 && player2[i-2] == 10 {
			score2 += player2[i]
		}
	}
	if score1 > score2 {
		return 1
	} else if score1 < score2 {
		return 2
	} else {
		return 0
	}
}

func main() {
	fmt.Println(isWinner([]int{5, 10, 3, 2}, []int{6, 5, 7, 3}))
}
