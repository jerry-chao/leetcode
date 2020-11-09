package main

import (
	"fmt"
	"leetcode/leetcode"
)

func main() {
	queue := leetcode.ConstructorMyCircularQueue(3)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	fmt.Println(queue.EnQueue(4))
	queue.Rear()
	queue.IsFull()
	queue.DeQueue()
	queue.EnQueue(4)
	queue.Rear()
}
