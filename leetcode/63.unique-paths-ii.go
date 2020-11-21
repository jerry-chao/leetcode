package leetcode

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 *
 * https://leetcode.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (34.88%)
 * Likes:    2207
 * Dislikes: 266
 * Total Accepted:    330.5K
 * Total Submissions: 945.7K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * Now consider if some obstacles are added to the grids. How many unique paths
 * would there be?
 *
 * An obstacle and space is marked as 1 and 0 respectively in the grid.
 *
 *
 * Example 1:
 *
 *
 * Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
 * Output: 2
 * Explanation: There is one obstacle in the middle of the 3x3 grid above.
 * There are two ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down -> Down
 * 2. Down -> Down -> Right -> Right
 *
 *
 * Example 2:
 *
 *
 * Input: obstacleGrid = [[0,1],[0,0]]
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * m == obstacleGrid.length
 * n == obstacleGrid[i].length
 * 1 <= m, n <= 100
 * obstacleGrid[i][j] is 0 or 1.
 *
 *
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			} else {
				if i > 0 && obstacleGrid[i-1][j] != -1 {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				}
				if j > 0 && obstacleGrid[i][j-1] != -1 {
					obstacleGrid[i][j] = obstacleGrid[i][j] + obstacleGrid[i][j-1]
				}
				if i == 0 && j == 0 {
					obstacleGrid[i][j] = 1
				}
			}
		}
	}
	result := obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
	if result < 0 {
		return 0
	}
	return result
}

// @lc code=end
