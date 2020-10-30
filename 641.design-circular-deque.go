package main

import "fmt"

/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */

// @lc code=start
type MyCircularDeque struct {
	Queue      []int
	Count      int
	Capacity   int
	FrontIndex int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	queue := make([]int, k)
	for i := 0; i < k; i++ {
		queue[i] = -1
	}
	return MyCircularDeque{
		Queue:    queue,
		Capacity: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.FrontIndex = (this.FrontIndex - 1 + this.Capacity) % this.Capacity
	this.Queue[this.FrontIndex] = value
	this.Count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.Queue[(this.FrontIndex+this.Count)%this.Capacity] = value
	this.Count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Queue[this.FrontIndex] = -1
	this.FrontIndex = (this.FrontIndex + 1) % this.Capacity
	this.Count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Queue[(this.FrontIndex+this.Count-1)%this.Capacity] = -1
	this.Count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	return this.Queue[this.FrontIndex]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Queue[(this.FrontIndex+this.Count-1)%this.Capacity]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.Count == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.Capacity == this.Count
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
// @lc code=end
func main() {

	obj := Constructor(3)
	param_1 := obj.InsertLast(1)
	param_2 := obj.InsertLast(2)
	param_3 := obj.InsertFront(3)
	param_4 := obj.InsertFront(4)
	fmt.Println(obj.FrontIndex, obj.Count)
	param_5 := obj.GetRear()
	param_6 := obj.IsFull()
	param_7 := obj.DeleteLast()
	param_8 := obj.InsertFront(4)
	param_9 := obj.GetFront()

	fmt.Println(param_1, param_2, param_3, param_4)
	fmt.Println(param_5, param_6, param_7, param_8)
	fmt.Println(param_9)
}
