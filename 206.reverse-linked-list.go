/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	p := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return p
// }

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre
}

// @lc code=end

