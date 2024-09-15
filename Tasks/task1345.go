package main

import (
	"fmt"
)

//func minJumps(arr []int) int {
//	n := len(arr)
//	memo := make(map[int]int)
//	var recursion func(index, steps int, visited map[int]bool) int
//	recursion = func(index, steps int, visited map[int]bool) int {
//		if index < 0 || index >= n || visited[index] {
//			return math.MaxInt32
//		}
//		if index == n-1 {
//			return steps
//		}
//		if val, ok := memo[index]; ok {
//			return val
//		}
//
//		visited[index] = true
//		minSteps := math.MaxInt32
//
//		// Переход вперед
//		minSteps = min(minSteps, recursion(index+1, steps+1, visited))
//
//		// Переход назад
//		minSteps = min(minSteps, recursion(index-1, steps+1, visited))
//
//		// Переход к другим индексам с одинаковыми значениями
//		for i := 0; i < n; i++ {
//			if i != index && arr[i] == arr[index] && !visited[i] {
//				minSteps = min(minSteps, recursion(i, steps+1, visited))
//			}
//		}
//
//		visited[index] = false
//		memo[index] = minSteps
//		return minSteps
//	}
//
//	return recursion(0, 0, make(map[int]bool))
//}

func minJumps(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}

	indexMap := make(map[int][]int)
	for i := 0; i < n; i++ {
		indexMap[arr[i]] = append(indexMap[arr[i]], i)
	}

	queue := [][]int{{0, 0}} // [index, steps]

	visited := make([]bool, n)
	visited[0] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		index, steps := curr[0], curr[1]

		if index == n-1 {
			return steps
		}

		if index+1 < n && !visited[index+1] {
			visited[index+1] = true
			queue = append(queue, []int{index + 1, steps + 1})
		}

		if index-1 >= 0 && !visited[index-1] {
			visited[index-1] = true
			queue = append(queue, []int{index - 1, steps + 1})
		}

		for _, nextIndex := range indexMap[arr[index]] {
			if nextIndex != index && !visited[nextIndex] {
				visited[nextIndex] = true
				queue = append(queue, []int{nextIndex, steps + 1})
			}
		}
		delete(indexMap, arr[index])
	}
	return -1
}

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}
