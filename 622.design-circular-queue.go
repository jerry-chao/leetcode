package main

import "fmt"

/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */

// @lc code=start
type MyCircularQueue struct {
	Queue      []int
	Capacity   int
	Count      int
	FrontIndex int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	queue := make([]int, k)
	for i := 0; i < k; i++ {
		queue[i] = -1
	}
	return MyCircularQueue{
		Queue:    queue,
		Capacity: k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Queue[(this.FrontIndex+this.Count)%this.Capacity] = value
	this.Count++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.Queue[this.FrontIndex] = -1
	this.Count--
	this.FrontIndex = (this.FrontIndex + 1) % this.Capacity
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	return this.Queue[this.FrontIndex]

}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[(this.FrontIndex+this.Count-1)%this.Capacity]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.Count == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.Count == this.Capacity
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

func main() {
	obj := Constructor(8)
	param_1 := obj.EnQueue(3)
	param_2 := obj.EnQueue(9)
	param_3 := obj.EnQueue(5)
	param_4 := obj.EnQueue(0)
	param_5 := obj.DeQueue()
	param_6 := obj.DeQueue()
	param_7 := obj.IsEmpty()
	param_8 := obj.IsEmpty()
	fmt.Println(obj.Count, obj.FrontIndex)
	param_9 := obj.Rear()
	param_10 := obj.Rear()
	param_11 := obj.DeQueue()

	fmt.Println(param_1, param_2, param_3, param_4, param_5, param_6, param_7)
	fmt.Println(param_8, param_9, param_10, param_11)
}
