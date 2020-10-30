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
	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < 2; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		tmp := tail.Next
		pre.Next = tail
		tail.Next = head
		head.Next = tmp
		tail = head
		pre = head
		head = tmp
	}
	return hair.Next
}

// @lc code=end
