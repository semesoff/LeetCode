package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type CharCount struct {
	Count int
	Char  byte
}

type MaxHeap []CharCount

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].Count > h[j].Count }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(CharCount)) }
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {
	pq := &MaxHeap{}

	if a > 0 {
		heap.Push(pq, CharCount{a, 'a'})
	}
	if b > 0 {
		heap.Push(pq, CharCount{b, 'b'})
	}
	if c > 0 {
		heap.Push(pq, CharCount{c, 'c'})
	}

	var res strings.Builder

	for pq.Len() > 0 {
		first := heap.Pop(pq).(CharCount)
		if res.Len() > 1 && res.String()[res.Len()-1] == first.Char && res.String()[res.Len()-2] == first.Char {
			if pq.Len() == 0 {
				break
			}
			second := heap.Pop(pq).(CharCount)
			res.WriteByte(second.Char)
			second.Count--
			if second.Count > 0 {
				heap.Push(pq, second)
			}
			heap.Push(pq, first)
		} else {
			res.WriteByte(first.Char)
			first.Count--
			if first.Count > 0 {
				heap.Push(pq, first)
			}
		}
	}
	return res.String()
}

func main() {
	fmt.Println(longestDiverseString(1, 1, 7)) // Output: "ccaccbcc"
}
