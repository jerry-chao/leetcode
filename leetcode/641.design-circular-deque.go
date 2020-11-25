package leetcode

/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */

// @lc code=start

// MyCircularDeque struce for circular deque
type MyCircularDeque struct {
	Queue      []int
	Count      int
	Capacity   int
	FrontIndex int
}

/*
ConstructorMyCircularDeque Initialize your data structure here. Set the size of the deque to be k.
*/
// func ConstructorMyCircularDeque(k int) MyCircularDeque {
func ConstructorMyCircularDeque(k int) MyCircularDeque {
	queue := make([]int, k)
	for i := 0; i < k; i++ {
		queue[i] = -1
	}
	return MyCircularDeque{
		Queue:    queue,
		Capacity: k,
	}
}

/*
InsertFront Adds an item at the front of Deque. Return true if the operation is successful.
*/
func (cd *MyCircularDeque) InsertFront(value int) bool {
	if cd.IsFull() {
		return false
	}
	frontIndex := (cd.FrontIndex - 1 + cd.Capacity) % cd.Capacity
	cd.Queue[frontIndex] = value
	cd.FrontIndex = frontIndex
	cd.Count++
	return true
}

/*
InsertLast Adds an item at the rear of Deque. Return true if the operation is successful.
*/
func (cd *MyCircularDeque) InsertLast(value int) bool {
	if cd.IsFull() {
		return false
	}
	lastIndex := (cd.FrontIndex + cd.Count) % cd.Capacity
	cd.Queue[lastIndex] = value
	cd.Count++
	return true
}

/*
DeleteFront Deletes an item from the front of Deque. Return true if the operation is successful.
*/
func (cd *MyCircularDeque) DeleteFront() bool {
	if cd.IsEmpty() {
		return false
	}
	cd.Queue[cd.FrontIndex] = -1
	cd.FrontIndex = (cd.FrontIndex + 1) % cd.Capacity
	cd.Count--
	return true
}

/*
DeleteLast Deletes an item from the rear of Deque. Return true if the operation is successful.
*/
func (cd *MyCircularDeque) DeleteLast() bool {
	if cd.IsEmpty() {
		return false
	}
	cd.Queue[(cd.FrontIndex+cd.Count-1)%cd.Capacity] = -1
	cd.Count--
	return true
}

/*
GetFront Get the front item from the deque.
*/
func (cd *MyCircularDeque) GetFront() int {
	if cd.IsEmpty() {
		return -1
	}
	return cd.Queue[cd.FrontIndex]
}

/*
GetRear Get the last item from the deque.
*/
func (cd *MyCircularDeque) GetRear() int {
	if cd.IsEmpty() {
		return -1
	}
	return cd.Queue[(cd.FrontIndex+cd.Count-1)%cd.Capacity]
}

/*
IsEmpty Checks whether the circular deque is empty or not.
*/
func (cd *MyCircularDeque) IsEmpty() bool {
	return cd.Count == 0
}

/*
IsFull Checks whether the circular deque is full or not.
*/
func (cd *MyCircularDeque) IsFull() bool {
	return cd.Capacity == cd.Count
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
