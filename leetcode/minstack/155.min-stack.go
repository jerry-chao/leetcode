package leetcode

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start

// MinStack minimue is the first element
type MinStack struct {
	Stack []int
	Min   []int
}

/*
Constructor initialize your data structure here.
*/
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}
}

// Push push x to min stack
func (m *MinStack) Push(x int) {
	m.Stack = append(m.Stack, x)
	minValue := x
	if len(m.Min) > 0 {
		minValue = min(m.GetMin(), x)
	}
	m.Min = append(m.Min, minValue)
}

// Pop pop the min value from stack
func (m *MinStack) Pop() {
	if len(m.Stack) > 0 {
		m.Stack = m.Stack[:len(m.Stack)-1]
		m.Min = m.Min[:len(m.Min)-1]
	}
}

// Top return top of min stack
func (m *MinStack) Top() int {
	if len(m.Stack) > 0 {
		return m.Stack[len(m.Stack)-1]
	}
	return -1
}

// GetMin get min of stack
func (m *MinStack) GetMin() int {
	if len(m.Stack) > 0 {
		return m.Min[len(m.Min)-1]
	}
	return -1
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
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
