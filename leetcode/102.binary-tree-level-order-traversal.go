package leetcode

import "log"

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (55.45%)
 * Likes:    3698
 * Dislikes: 88
 * Total Accepted:    704.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 *
 *
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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return levelOrderLocal(root, 0, [][]int{})
}

func levelOrderLocal(root *TreeNode, dep int, result [][]int) [][]int {
	// terminator
	if root == nil {
		return result
	}
	// handle
	if len(result) > dep {
		result[dep] = append(result[dep], root.Val)
	} else {
		result = append(result, []int{root.Val})
	}
	result = levelOrderLocal(root.Left, dep+1, result)
	log.Println(result, root.Val, dep)
	return levelOrderLocal(root.Right, dep+1, result)
}

// @lc code=end
