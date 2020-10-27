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
	if head == nil {
		return false
	}
	slow := head
	fast := head
	steps := 0
	for fast != nil && (fast != slow || steps == 0) {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
		steps++
	}
	fmt.Println(fast, slow, steps)
	if fast == slow && steps != 0 {
		return true
	}
	return false
}

// @lc code=end

