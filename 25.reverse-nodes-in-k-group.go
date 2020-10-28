package main

/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		tmp := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		pre = tail
		head = tmp
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	current := head
	for pre != tail {
		tmp := current.Next
		current.Next = pre
		pre = current
		current = tmp
	}
	return tail, head
}

// @lc code=end
