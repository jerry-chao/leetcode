package leetcode

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

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	// the last node of reversed link list
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		// store tail.Next
		tmp := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		pre = tail
		head = tmp
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	// store pre head linked list
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

// func main() {
// 	Node5 := &ListNode{Val: 5}
// 	Node4 := &ListNode{Val: 4, Next: Node5}
// 	Node3 := &ListNode{Val: 3, Next: Node4}
// 	Node2 := &ListNode{Val: 2, Next: Node3}
// 	Node1 := &ListNode{Val: 1, Next: Node2}
// 	Node := reverseKGroup(Node1, 2)
// 	for Node != nil {
// 		fmt.Print(Node.Val, " ")
// 		Node = Node.Next
// 	}
// }
