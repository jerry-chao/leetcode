package main

import "fmt"

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	Stack []int
	Min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	// push x to stack
	this.Stack = append(this.Stack, x)
	// min
	if len(this.Min) > 0 {
		this.Min = append(this.Min, min(this.GetMin(), x))
	} else {
		this.Min = append(this.Min, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) > 0 {
		this.Stack = this.Stack[:len(this.Stack)-1]
		this.Min = this.Min[:len(this.Min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) > 0 {
		return this.Stack[len(this.Stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.Min) > 0 {
		return this.Min[len(this.Min)-1]
	}
	return -1
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param_1 := obj.GetMin()
	obj.Pop()
	param_2 := obj.Top()
	param_3 := obj.GetMin()
	fmt.Println(param_1, param_2, param_3)
}
