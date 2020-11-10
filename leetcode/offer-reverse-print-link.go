package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	return reversePrintRec(head, []int{})
}
func reversePrintRec(head *ListNode, result []int) []int {
	// terminator
	if head == nil {
		return result
	}
	// current
	result = append([]int{head.Val}, result...)
	// drill down
	return reversePrintRec(head.Next, result)
}
