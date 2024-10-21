package main

func countNeighbours(board [][]int, i, j int) int {
	r := len(board)
	c := len(board[0])
	count := 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x >= 0 && x < r && y >= 0 && y < c && (x != i || y != j) && board[x][y] == 1 {
				count++
			}
		}
	}
	return count
}

func gameOfLife(board [][]int) {
	r := len(board)
	c := len(board[0])
	preGen := make([][]int, r)
	for i := range board {
		preGen[i] = append([]int{}, board[i]...)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			cn := countNeighbours(preGen, i, j)
			if preGen[i][j] == 1 {
				if cn < 2 || cn > 3 {
					board[i][j] = 0
				}
			} else {
				if cn == 3 {
					board[i][j] = 1
				}
			}
		}
	}
}

func main() {
	gameOfLife([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}})
}
