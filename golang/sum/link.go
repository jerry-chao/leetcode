package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println(10 / 10)
	fmt.Println(10 % 10)
	l1 := makeList([]int{2, 4, 3})
	l2 := makeList([]int{5, 6, 4})
	printListNode(addTwoNumbers(l1, l2))
	printListNode(addTwoNumbers(makeList([]int{1, 8}), makeList([]int{0})))
	printListNode(addTwoNumbers(makeList([]int{5}), makeList([]int{5})))
}

//  printListNode ...
func printListNode(l *ListNode) {
	for l != nil {
		fmt.Println("this is list node:", l.Val)
		l = l.Next
	}
}

func makeList(l []int) *ListNode {
	var ret *ListNode
	var tmp *ListNode
	for _, value := range l {
		if ret == nil {
			ret = &ListNode{Val: value}
			tmp = ret
		} else {
			tmp.Next = &ListNode{Val: value}
			tmp = tmp.Next
		}
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	inc := 0
	val1 := 0
	val2 := 0
	var tmp, head *ListNode
	for l1 != nil || l2 != nil {
		val1 = getVal(l1)
		val2 = getVal(l2)
		fmt.Println("xxxx", val1, val2, inc)
		if tmp == nil {
			tmp = &ListNode{Val: (val1 + val2 + inc) % 10}
			head = tmp
		} else {
			tmp.Next = &ListNode{Val: (val1 + val2 + inc) % 10}
			tmp = tmp.Next
		}
		inc = (val1 + val2 + inc) / 10
		l1 = moveNext(l1)
		l2 = moveNext(l2)
	}

	if inc == 0 {
		return head
	}

	tmp.Next = &ListNode{Val: inc}

	return head
}

func moveNext(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func reverse(l *ListNode) *ListNode {
	var ret *ListNode
	var tmp *ListNode
	for l != nil {
		tmp = l
		tmp.Next = ret
		ret = tmp
		l = l.Next
	}
	return ret
}
