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

func largestValues(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	resultLargestValues = map[int]int{}
	visLargestValues = map[*TreeNode]bool{}
	dfsLargestValues(root, 0)
	// handle resultMap
	for i := 0; i < len(resultLargestValues); i++ {
		result = append(result, resultLargestValues[i])
	}
	return result
}

// dfs
var resultLargestValues map[int]int
var visLargestValues map[*TreeNode]bool

func dfsLargestValues(root *TreeNode, dep int) {
	if root == nil {
		return
	}
	if _, ok := visLargestValues[root]; ok {
		return
	}
	if old, ok := resultLargestValues[dep]; ok {
		if old < root.Val {
			resultLargestValues[dep] = root.Val
		}
	} else {
		resultLargestValues[dep] = root.Val
	}
	// drill down
	dfsLargestValues(root.Left, dep+1)
	dfsLargestValues(root.Right, dep+1)
}

// @lc code=end
