package leetcode

import (
	"container/heap"
)

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start

// TopKFrequent return top k frequent element
func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

func topKFrequent(nums []int, k int) []int {
	result := []int{}
	if len(nums) < k {
		return result
	}
	// key freq
	frequents := map[int]int{}
	for i := 0; i < len(nums); i++ {
		frequents[nums[i]]++
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	for fKey, fValue := range frequents {
		heap.Push(pq,
			&Item{value: fKey,
				priority: fValue})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	// relate result
	for i := 0; i < k; i++ {
		item := heap.Pop(pq).(*Item)
		result = append(result, item.value)
	}
	return result
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int
	priority int
}

// PriorityQueue return element with priority
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push push x into PriorityQueue
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop pop priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// @lc code=end
