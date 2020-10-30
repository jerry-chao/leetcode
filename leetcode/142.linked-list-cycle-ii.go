package leetcode

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	first, second := head, head
	step := 0
	for first != second || step == 0 {
		if second == nil {
			return nil
		}
		first = first.Next
		second = second.Next
		if second == nil {
			return nil
		}
		second = second.Next
		step++
	}
	current := head
	for current != second {
		current = current.Next
		second = second.Next
	}
	return current
}

// @lc code=end
