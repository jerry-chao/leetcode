package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 *
 * https://leetcode.com/problems/triangle/description/
 *
 * algorithms
 * Medium (44.85%)
 * Likes:    2423
 * Dislikes: 271
 * Total Accepted:    271K
 * Total Submissions: 601.5K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * Given a triangle, find the minimum path sum from top to bottom. Each step
 * you may move to adjacent numbers on the row below.
 *
 * For example, given the following triangle
 *
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
 *
 * Note:
 *
 * Bonus point if you are able to do this using only O(n) extra space, where n
 * is the total number of rows in the triangle.
 *
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// DP[row][col] = min(DP[row][col], DP[row][col-1]) + triangle[row][col]
	if len(triangle) == 0 {
		return 0
	}
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] = triangle[i-1][0] + triangle[i][0]
		for j := 1; j < len(triangle[i]); j++ {
			if j < len(triangle[i-1]) {
				triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j-1]) + triangle[i][j]
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			}
		}
	}
	minTotal := math.MaxInt64
	lastRow := triangle[len(triangle)-1]
	for i := 0; i < len(lastRow); i++ {
		minTotal = min(lastRow[i], minTotal)
	}
	return minTotal
}

// func min(i, j int) int {
// 	if i > j {
// 		return j
// 	}
// 	return i
// }

// @lc code=end
