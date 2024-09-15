package main

import (
	"fmt"
)

func createMap(graphAlice *map[int]map[int]int, graphBob *map[int]map[int]int, edges [][]int) int {
	need := 0
	for _, j := range edges {
		if j[0] == 3 {
			need++
			if _, ok := (*graphAlice)[j[1]]; !ok {
				(*graphAlice)[j[1]] = map[int]int{}
			}
			if _, ok := (*graphAlice)[j[2]]; !ok {
				(*graphAlice)[j[2]] = map[int]int{}
			}
			if _, ok := (*graphBob)[j[1]]; !ok {
				(*graphBob)[j[1]] = map[int]int{}
			}
			if _, ok := (*graphBob)[j[2]]; !ok {
				(*graphBob)[j[2]] = map[int]int{}
			}
			(*graphAlice)[j[1]][j[2]] = 1
			(*graphAlice)[j[2]][j[1]] = 1
			(*graphBob)[j[1]][j[2]] = 1
			(*graphBob)[j[2]][j[1]] = 1

		}
	}
	return need
}

func addValue(graph *map[int]map[int]int, val1, val2 int) {
	if _, ok := (*graph)[val1]; !ok {
		(*graph)[val1] = map[int]int{}
	}
	if _, ok := (*graph)[val2]; !ok {
		(*graph)[val2] = map[int]int{}
	}
	(*graph)[val1][val2] = 1
	(*graph)[val2][val1] = 1
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	graphAlice := make(map[int]map[int]int)
	graphBob := make(map[int]map[int]int)
	edgesNeeded := createMap(&graphAlice, &graphBob, edges)
	edgesTotal := edgesNeeded
	for _, i := range edges {
		if i[0] == 1 {
			if _, ok := graphAlice[i[1]]; ok {
				if _, ok1 := graphAlice[i[2]]; ok1 {
					addValue(&graphAlice, i[1], i[2])
					edgesNeeded++
				}

			}
			edgesTotal++
		} else if i[0] == 2 {
			if _, ok := graphBob[i[1]]; ok {
				if _, ok1 := graphBob[i[2]]; ok1 {
					addValue(&graphBob, i[1], i[2])
					edgesNeeded++
				}
			}
			edgesTotal++
		}
	}
	return edgesTotal - edgesNeeded
}

func main() {
	fmt.Println(maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}}))
}
