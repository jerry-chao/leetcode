package leetcode

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
	if root == nil {
		return nil
	}
	// left
	left := inorderTraversal(root.Left)
	result := append(left, root.Val)
	// right
	right := inorderTraversal(root.Right)
	return append(result, right...)
}

// @lc code=end
