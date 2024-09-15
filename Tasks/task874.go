package main

import (
	"fmt"
	"math"
)

func movement(x, y *int, steps int, vector int, obstacles map[[2]int]int) {
	oldCord := [2]int{0, 0}
	for i := 0; i < steps; i++ {
		oldCord = [2]int{*x, *y}
		if vector == 1 {
			*y++
		} else if vector == 2 {
			*x++
		} else if vector == 3 {
			*y--
		} else if vector == 4 {
			*x--
		}
		if _, ok := obstacles[[2]int{*x, *y}]; ok {
			*x = oldCord[0]
			*y = oldCord[1]
			break
		}
	}
}

func action(x, y *int, act int, vector *int, obstacles map[[2]int]int) {
	if act >= 1 && act <= 9 {
		movement(x, y, act, *vector, obstacles)
	} else if act == -1 || act == -2 {
		changeVector(vector, act)
	}
}

func changeVector(vector *int, act int) {
	if act == -2 {
		*vector--
	} else {
		*vector++
	}
	if *vector == 0 {
		*vector = 4
	} else if *vector == 5 {
		*vector = 1
	}
}

func robotSim(commands []int, obstacles [][]int) int {
	hash := make(map[[2]int]int)
	for _, val := range obstacles {
		hash[[2]int(val)] = 1
	}
	x, y := 0, 0
	maxDistance := 0
	vector := 1
	for _, command := range commands {
		action(&x, &y, command, &vector, hash)
		maxDistance = int(math.Max(float64(maxDistance), float64(x*x+y*y)))
	}
	return maxDistance
}

func main() {
	fmt.Println(robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}))
}
