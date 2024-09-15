package main

import (
	"container/heap"
	"fmt"
)

type Project struct {
	profit   int
	capital  int
}


type MinCapitalHeap []Project

func (h MinCapitalHeap) Len() int           { return len(h) }
func (h MinCapitalHeap) Less(i, j int) bool { return h[i].capital < h[j].capital }
func (h MinCapitalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinCapitalHeap) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MinCapitalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxProfitHeap реализует максимальную кучу для прибыли.
type MaxProfitHeap []Project

func (h MaxProfitHeap) Len() int           { return len(h) }
func (h MaxProfitHeap) Less(i, j int) bool { return h[i].profit > h[j].profit }
func (h MaxProfitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxProfitHeap) Push(x interface{}) {
	*h = append(*h, x.(Project))
}

func (h *MaxProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// findMaximizedCapital реализует основную логику решения.
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minCapitalHeap := &MinCapitalHeap{}
	maxProfitHeap := &MaxProfitHeap{}

	// Заполнить минимальную кучу капиталов
	for i := 0; i < len(profits); i++ {
		heap.Push(minCapitalHeap, Project{profit: profits[i], capital: capital[i]})
	}

	for i := 0; i < k; i++ {
		// Перемещать проекты, которые можно выполнить, в кучу прибыли
		for minCapitalHeap.Len() > 0 && (*minCapitalHeap)[0].capital <= w {
			heap.Push(maxProfitHeap, heap.Pop(minCapitalHeap).(Project))
		}

		// Если нет доступных проектов, завершить
		if maxProfitHeap.Len() == 0 {
			break
		}

		// Выполнить проект с наибольшей прибылью
		w += heap.Pop(maxProfitHeap).(Project).profit
	}

	return w
}

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}

	fmt.Println(findMaximizedCapital(k, w, profits, capital)) // Output: 4
}
