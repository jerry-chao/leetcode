package leetcode

/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (55.45%)
 * Likes:    3718
 * Dislikes: 88
 * Total Accepted:    706.9K
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
func levelOrder(root *TreeNode) [][]int {
	// BFS
	queueLevelOrder := []*TreeNode{root}
	result := [][]int{}
	for len(queueLevelOrder) > 0 {
		levelResult := []int{}
		levelQueue := queueLevelOrder
		queueLevelOrder = []*TreeNode{}
		for len(levelQueue) > 0 {
			// pop first value
			top := levelQueue[0]
			if top != nil {
				levelResult = append(levelResult, top.Val)
				queueLevelOrder = append(queueLevelOrder, top.Left, top.Right)
			}
			// log.Println(top)
			if len(levelQueue) > 1 {
				levelQueue = levelQueue[1:]
			} else {
				levelQueue = []*TreeNode{}
			}
		}
		if len(levelResult) > 0 {
			result = append(result, levelResult)
		}
	}
	return result
}

// @lc code=end
