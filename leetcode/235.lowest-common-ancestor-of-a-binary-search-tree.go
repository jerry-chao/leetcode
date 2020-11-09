package leetcode

/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestorBinaryTree(root, p, q *TreeNode) *TreeNode {
	// terminator
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// handle current process
	// drill down
	left := lowestCommonAncestorBinaryTree(root.Left, p, q)
	right := lowestCommonAncestorBinaryTree(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// @lc code=end
