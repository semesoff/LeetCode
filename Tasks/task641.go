package main

import "fmt"

type MyCircularDeque struct {
	deque      []int
	capacity   int
	size       int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		deque:      make([]int, 0),
		capacity:   k,
		size:       0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.size == this.capacity {
		return false
	}
	this.deque = append([]int{value}, this.deque...)
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.size == this.capacity {
		return false
	}
	this.deque = append(this.deque, value)
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.size == 0 {
		return false
	}
	this.deque = this.deque[1:]
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.size == 0 {
		return false
	}
	this.deque = this.deque[:len(this.deque)-1]
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.size == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.size == 0 {
		return -1
	}
	return this.deque[this.size-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.capacity
}

func main() {
	circularDeque := NewMyCircularDeque(3)
	fmt.Println(circularDeque.InsertLast(1))  // return true
	fmt.Println(circularDeque.InsertLast(2))  // return true
	fmt.Println(circularDeque.InsertFront(3)) // return true
	fmt.Println(circularDeque.InsertFront(4)) // return false, the queue is full
	fmt.Println(circularDeque.GetRear())      // return 2
	fmt.Println(circularDeque.IsFull())       // return true
	fmt.Println(circularDeque.DeleteLast())   // return true
	fmt.Println(circularDeque.InsertFront(4)) // return true
	fmt.Println(circularDeque.GetFront())     // return 4
	fmt.Println(circularDeque)                // return 2
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
