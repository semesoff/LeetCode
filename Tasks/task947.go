package main

import "fmt"

func removeStones(stones [][]int) int {
	connectedComponentCount := 0
	setRepresentatives := make(map[int]int)

	var findRepresentative func(int) int
	findRepresentative = func(element int) int {
		if _, exists := setRepresentatives[element]; !exists {
			setRepresentatives[element] = element
			connectedComponentCount++
		}
		if setRepresentatives[element] != element {
			setRepresentatives[element] = findRepresentative(setRepresentatives[element])
		}
		return setRepresentatives[element]
	}

	mergeComponents := func(elementA, elementB int) {
		repA := findRepresentative(elementA)
		repB := findRepresentative(elementB)
		if repA != repB {
			setRepresentatives[repB] = repA
			connectedComponentCount--
		}
	}

	for _, stone := range stones {
		mergeComponents(stone[0]+1, stone[1]+10002)
	}

	return len(stones) - connectedComponentCount
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}))
}
