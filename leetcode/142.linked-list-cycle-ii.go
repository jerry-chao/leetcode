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

	// find cycle
	slow := head
	fast := head
	steps := 0
	for fast != nil && (slow != fast || steps == 0) {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
		steps++
	}
	if slow == fast && steps != 0 {
		// there are cycle
		for head != slow {
			head = head.Next
			slow = slow.Next
		}
		return head
	}
	return nil
}

// @lc code=end
