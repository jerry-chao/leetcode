package main

import (
	"fmt"
)

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	fmt.Println("hello")
	fmt.Println(ListNode{val: 1})
	a := ListNode{val: 1}
	b := ListNode{val: 2, next: &a}
	c := ListNode{val: 3, next: &b}

	e := ListNode{val: 4, next: &b}
	f := ListNode{val: 5, next: &e}
	g := ListNode{val: 6, next: &f}
	fmt.Println(getIntersectionNode(&g, &c))
	// printListNode(&g)
	// printListNode(reverseListNode(&g))

	// printListNode(&c)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headA = reverseListNode(headA)
	printListNode(headA)

	fmt.Println("xxxxxxxxx")
	printListNode(headB)
	headB = reverseListNode(headB)
	fmt.Println(headA)
	fmt.Println(headB)
	var first *ListNode
	for headA.val == headB.val && headA != nil && headB != nil {
		first = headA
		headA = headA.next
		headB = headB.next
	}
	return first
}

func reverseListNode(head *ListNode) *ListNode {
	var tmp, list *ListNode
	for head != nil {
		tmp = &ListNode{val: head.val, next: list}
		list = tmp
		head = head.next
	}
	return list
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Println("=====>", head.val)
		head = head.next
	}
}
