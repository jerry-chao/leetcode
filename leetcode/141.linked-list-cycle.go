package leetcode

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// first move forward 1 step and second move 2 step
	first, second := head, head
	step := 0
	for first != second || step == 0 {
		if second == nil {
			return false
		}
		first = first.Next
		second = second.Next
		if second == nil {
			return false
		}
		step++
		second = second.Next
	}
	return true
}

// @lc code=end
