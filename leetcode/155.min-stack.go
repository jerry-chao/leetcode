package leetcode

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
	this.Stack = append(this.Stack, x)
	minValue := x
	if len(this.Min) > 0 {
		minValue = min(this.GetMin(), x)
	}
	this.Min = append(this.Min, minValue)
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
	} else {
		return -1
	}
}

func (this *MinStack) GetMin() int {
	if len(this.Stack) > 0 {
		return this.Min[len(this.Min)-1]
	} else {
		return -1
	}
}

// func min(i, j int) int {
// 	if i > j {
// 		return j
// 	}
// 	return i
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
