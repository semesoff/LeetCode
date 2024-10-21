package main

import "fmt"

func isValidBox(board [][]byte, i, j int) bool {
	m := make(map[byte]bool)
	for x := i; x < i+3; x++ {
		for y := j; y < j+3; y++ {
			if board[x][y] != '.' {
				if m[board[x][y]] {
					return false
				}
				m[board[x][y]] = true
			}
		}
	}
	return true
}

func isValidRow(board [][]byte, x int) bool {
	for i := x; i < x+3; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if m[board[i][j]] {
					return false
				}
				m[board[i][j]] = true
			}
		}
	}
	return true
}

func isValidCol(board [][]byte, x int) bool {
	for j := x; j < x+3; j++ {
		m := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if m[board[i][j]] {
					return false
				}
				m[board[i][j]] = true
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i += 3 {
		if !isValidRow(board, i) || !isValidCol(board, i) {
			return false
		}
		for j := 0; j < 9; j += 3 {
			if !isValidBox(board, i, j) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isValidSudoku([][]byte{}))
}
