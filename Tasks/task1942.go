package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	friends := make(map[[2]int]int, n)
	for i, time := range times {
		friends[[2]int{time[0], time[1]}] = i
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i][0] < times[j][0]
	})

	available := &MinHeap{}
	heap.Init(available)
	for i := 0; i < n; i++ {
		heap.Push(available, i)
	}

	occupied := make([][2]int, 0)

	for _, time := range times {
		currentTime := time[0]
		for len(occupied) > 0 && occupied[0][0] <= currentTime {
			heap.Push(available, occupied[0][1])
			occupied = occupied[1:]
		}

		chair := heap.Pop(available).(int)
		occupied = append(occupied, [2]int{time[1], chair})
		sort.Slice(occupied, func(i, j int) bool {
			return occupied[i][0] < occupied[j][0]
		})

		if friends[[2]int{time[0], time[1]}] == targetFriend {
			return chair
		}
	}
	return -1
}

func main() {
	fmt.Println(smallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0))
}
