package main

import (
	"fmt"
	"maps"
	"math"
)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	hash := make(map[int][]int)
	scores := make(map[int]map[int]float64)
	for i, val := range edges {
		hash[val[0]] = append(hash[val[0]], val[1])
		hash[val[1]] = append(hash[val[1]], val[0])
		if _, ok := scores[val[0]]; !ok {
			scores[val[0]] = make(map[int]float64)
		}
		if _, ok := scores[val[1]]; !ok {
			scores[val[1]] = make(map[int]float64)
		}
		scores[val[0]][val[1]] = succProb[i]
		scores[val[1]][val[0]] = succProb[i]
	}

	var findPath func(current int, path float64, visited map[int]int) float64
	findPath = func(current int, path float64, visited map[int]int) float64 {
		if _, ok := visited[current]; ok {
			return 0
		}
		if current == end_node {
			return path
		}
		visited[current] = 1
		maxPath := float64(0)
		for _, i := range hash[current] {
			maxPath = math.Max(maxPath, findPath(i, path*scores[current][i], maps.Clone(visited)))
		}
		return maxPath
	}

	return findPath(start_node, 1, make(map[int]int))
}

func main() {
	fmt.Println(maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
}
