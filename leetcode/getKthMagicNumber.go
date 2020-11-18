package leetcode

import "container/heap"

func getKthMagicNumber(k int) int {
	// BFS
	magic := []int{3, 5, 7}
	queue := &PriorityQueue{}
	heap.Init(queue)
	heap.Push(queue, &Item{value: 1, priority: 1})
	vis := map[int]int{}
	for true {
		// pop queue
		ele := heap.Pop(queue).(*Item).value
		if _, ok := vis[ele]; !ok {
			vis[ele] = 1
			for i := 0; i < len(magic); i++ {
				newEle := ele * magic[i]
				heap.Push(queue, &Item{value: newEle, priority: newEle})
			}
		}
		if len(vis) == k {
			return ele
		}
	}
	return -1
}
