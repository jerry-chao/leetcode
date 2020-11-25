package leetcode

import "fmt"

// TreeNode tree node struct binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode list node struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// Node for tree node
type Node struct {
	Val      int
	Children []*Node
}

// PrintArray print array for int
func PrintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("")
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
