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

// func topKFrequent(nums []int, k int) []int {
// 	numsWithFrequent := map[int]int{}
// 	for i := 0; i < len(nums); i++ {
// 		numsWithFrequent[nums[i]]++
// 	}
// 	if len(numsWithFrequent) < k {
// 		return []int{}
// 	}

// 	pq := PriorityQueue{}
// 	for key, value := range numsWithFrequent {
// 		pq.Push(&Item{
// 			value:    key,
// 			priority: value,
// 		})
// 		if pq.Len() > k {
// 			pq.Pop()
// 		}
// 	}
// 	result := make([]int, k)
// 	for i := k - 1; i >= 0; i-- {
// 		item := pq.Pop().(*Item)
// 		result[i] = item.value
// 	}
// 	return result
// }

// type Item struct {
// 	value    int // The value of the item; arbitrary.
// 	priority int // The priority of the item in the queue.
// }

// // A PriorityQueue implements heap.Interface and holds Items.
// type PriorityQueue []*Item

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
// 	return pq[i].priority > pq[j].priority
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// func (pq *PriorityQueue) Push(x interface{}) {
// 	fmt.Println("push item:", x)
// 	item := x.(*Item)
// 	*pq = append(*pq, item)
// }

// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	fmt.Println("pop item:", item)
// 	old[n-1] = nil // avoid memory leak
// 	*pq = old[:n-1]
// 	return item
// }

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
