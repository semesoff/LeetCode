package main

import "fmt"

type Stack [][2]int

func (s *Stack) add(value [2]int) {
	*s = append(*s, value)
}

func (s *Stack) pop() (int, int) {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item[0], item[1]
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func countIslands(n, m int, grid [][]int) int {
	seen := make(map[[2]int]int)
	var dfs func(r, c int)
	dfs = func(r, c int) {
		stack := Stack{[2]int{r, c}}
		for !stack.isEmpty() {
			x, y := stack.pop()
			for _, val := range [4][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
				nx := val[0]
				ny := val[1]
				if _, ok := seen[[2]int{nx, ny}]; 0 <= nx && nx < n && 0 <= ny && ny < m && !ok && grid[nx][ny] == 1 {
					seen[[2]int{nx, ny}] = 1
					stack.add([2]int{nx, ny})
				}
			}
		}
	}

	islands := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := seen[[2]int{i, j}]; grid[i][j] == 1 && !ok {
				islands++
				seen[[2]int{i, j}] = 1
				dfs(i, j)
			}
		}
	}
	return islands
}

func minDays(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	// if there are no islands and there are two islands
	if countIslands(n, m, grid) != 1 {
		return 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				if countIslands(n, m, grid) != 1 {
					return 1
				}
				grid[i][j] = 1
			}
		}
	}
	return 2
}

func main() {
	fmt.Println(minDays([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
}
