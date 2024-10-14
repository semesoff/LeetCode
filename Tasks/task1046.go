package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}
	fmt.Println(maxHeap)
	for maxHeap.Len() > 1 {
		x := heap.Pop(maxHeap).(int)
		y := heap.Pop(maxHeap).(int)
		if x != y {
			heap.Push(maxHeap, x-y)
		}
	}
	if maxHeap.Len() != 0 {
		return maxHeap.Pop().(int)
	}
	return 0
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
