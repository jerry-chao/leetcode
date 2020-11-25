package leetcode

/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */

// @lc code=start

// MyCircularQueue design for circular queue
type MyCircularQueue struct {
	Queue      []int
	Capacity   int
	Count      int
	FrontIndex int
}

// ConstructorMyCircularQueue init my circular queue
func ConstructorMyCircularQueue(k int) MyCircularQueue {
	queue := make([]int, k)
	for i := 0; i < k; i++ {
		queue[i] = -1
	}
	return MyCircularQueue{
		Queue:    queue,
		Capacity: k,
	}
}

/*
EnQueue Insert an element into the circular queue. Return true if the operation is successful.
*/
func (cq *MyCircularQueue) EnQueue(value int) bool {
	if cq.IsFull() {
		return false
	}
	index := (cq.FrontIndex + cq.Count) % cq.Capacity
	cq.Queue[index] = value
	cq.Count++
	return true
}

/*
DeQueue Delete an element from the circular queue. Return true if the operation is successful.
*/
func (cq *MyCircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		return false
	}
	cq.Queue[cq.FrontIndex] = -1
	cq.FrontIndex = (cq.FrontIndex + 1) % cq.Capacity
	cq.Count--
	return true
}

/*
Front Get the front item from the queue.
*/
func (cq *MyCircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.Queue[cq.FrontIndex]
}

/*
Rear Get the last item from the queue.
*/
func (cq *MyCircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}
	lastIndex := (cq.FrontIndex + cq.Count - 1) % cq.Capacity
	return cq.Queue[lastIndex]
}

/*
IsEmpty Checks whether the circular queue is empty or not.
*/
func (cq *MyCircularQueue) IsEmpty() bool {
	return cq.Count == 0
}

/*
IsFull Checks whether the circular queue is full or not.
*/
func (cq *MyCircularQueue) IsFull() bool {
	return cq.Count == cq.Capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
