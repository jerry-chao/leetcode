package main

import "fmt"

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{}
	// handle root left node
	current := root
	for current != nil {
		stack = append(stack, current)
		current = current.Left
	}

	for len(stack) > 0 {
		// pop last element
		lastElment := stack[len(stack)-1]
		// handle root
		result = append(result, lastElment.Val)
		stack = stack[:len(stack)-1]
		// handle right sub tree
		current = lastElment.Right
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
	}

	return result
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	Node3 := &TreeNode{Val: 3}
	Node2 := &TreeNode{Val: 2, Left: Node3}
	Node := &TreeNode{Val: 1, Right: Node2}
	for _, num := range inorderTraversal(Node) {
		fmt.Print(num, " ")
	}

}
