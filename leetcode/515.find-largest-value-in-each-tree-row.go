package leetcode

// @lc code=start
// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * } */

//  */

var maxLargestValues map[int]int
var visitedLargestValues map[*TreeNode]bool

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	maxLargestValues = map[int]int{}
	visitedLargestValues = map[*TreeNode]bool{}
	dfsLargestValues(0, root)
	result := []int{}
	for i := 0; i < len(maxLargestValues); i++ {
		result = append(result, maxLargestValues[i])
	}
	return result
}

func dfsLargestValues(dep int, node *TreeNode) {
	if node == nil {
		return
	}
	if had, ok := maxLargestValues[dep]; ok {
		if had < node.Val {
			maxLargestValues[dep] = node.Val
		}
	} else {
		maxLargestValues[dep] = node.Val
	}
	dfsLargestValues(dep+1, node.Left)
	dfsLargestValues(dep+1, node.Right)
}

// @lc code=end
