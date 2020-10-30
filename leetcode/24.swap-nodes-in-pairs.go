package leetcode

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	first := head
	if first.Next == nil {
		return head
	}
	second := first.Next
	current := &ListNode{}
	head = current
	for second != nil {
		current.Next = second
		tmp := second.Next
		second.Next = first
		first.Next = tmp
		current = first
		if tmp == nil {
			break
		}
		first = tmp
		second = first.Next
	}
	return head.Next
}

// @lc code=end
