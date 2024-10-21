package main

type MinStack struct {
	Stack []int
	Size  int
	Min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.Min) == 0 || val <= this.Min[len(this.Min)-1] {
		this.Min = append(this.Min, val)
	}
	this.Size++
}

func (this *MinStack) Pop() {
	if this.Stack[this.Size-1] == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
	this.Stack = this.Stack[:this.Size-1]
	this.Size--
}

func (this *MinStack) Top() int {
	return this.Stack[this.Size-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
