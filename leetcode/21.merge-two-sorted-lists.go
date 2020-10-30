package main

import "fmt"

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return head.Next
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	Node13 := &ListNode{Val: 4}
	Node12 := &ListNode{Val: 2, Next: Node13}
	Node11 := &ListNode{Val: 1, Next: Node12}
	Node23 := &ListNode{Val: 4}
	Node22 := &ListNode{Val: 3, Next: Node23}
	Node21 := &ListNode{Val: 1, Next: Node22}
	Node1 := mergeTwoLists(Node11, Node21)
	printLinkList(Node1)
}

func printLinkList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
