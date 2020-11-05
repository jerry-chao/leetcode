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
func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}
	pq := &PriorityQueue{}
	heap.Init(pq)
	for key, value := range hash {
		heap.Push(pq, &Item{
			value:    key,
			priority: value,
		})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	result := []int{}
	for pq.Len() > 0 {
		item := pq.Pop().(*Item)
		result = append(result, item.value)
	}
	return result
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return item
}

// @lc code=end
