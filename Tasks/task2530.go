package main

import (
	"container/heap"
	"fmt"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) PopLeft() interface{} {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}

func (h *IntHeap) ChangeValue(i, x int) {
	(*h)[i] = x
	heap.Fix(h, i)
}

func (h *IntHeap) Peek() interface{} {
	return (*h)[0]
}

func maxKelements(nums []int, k int) int64 {
	maxHeap := &IntHeap{}
	for _, num := range nums {
		heap.Push(maxHeap, num)
	}
	heap.Init(maxHeap)

	res := 0
	for k > 0 {
		x := maxHeap.Peek().(int)
		res += x
		k--
		maxHeap.ChangeValue(0, int(math.Ceil(float64(x)/3)))
	}
	return int64(res)
}

func main() {
	fmt.Println(maxKelements([]int{1, 10, 3, 3, 3}, 3)) // 12
}
