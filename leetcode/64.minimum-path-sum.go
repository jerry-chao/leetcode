package leetcode

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (55.23%)
 * Likes:    3830
 * Dislikes: 75
 * Total Accepted:    475.7K
 * Total Submissions: 858.2K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right, which minimizes the sum of all numbers along its
 * path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
 * Output: 7
 * Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[1,2,3],[4,5,6]]
 * Output: 12
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 100
 *
 *
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + nums[i][j]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 && j > 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			} else if i > 0 && j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if j > 0 && i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

// func min(i, j int) int {
// 	if i > j {
// 		return j
// 	}
// 	return i
// }

// @lc code=end
