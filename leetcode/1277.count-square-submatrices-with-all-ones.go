package leetcode

/*
 * @lc app=leetcode id=1277 lang=golang
 *
 * [1277] Count Square Submatrices with All Ones
 *
 * https://leetcode.com/problems/count-square-submatrices-with-all-ones/description/
 *
 * algorithms
 * Medium (73.06%)
 * Likes:    1403
 * Dislikes: 25
 * Total Accepted:    78.5K
 * Total Submissions: 107.5K
 * Testcase Example:  '[[0,1,1,1],[1,1,1,1],[0,1,1,1]]'
 *
 * Given a m * n matrix of ones and zeros, return how many square submatrices
 * have all ones.
 *
 *
 * Example 1:
 *
 *
 * Input: matrix =
 * [
 * [0,1,1,1],
 * [1,1,1,1],
 * [0,1,1,1]
 * ]
 * Output: 15
 * Explanation:
 * There are 10 squares of side 1.
 * There are 4 squares of side 2.
 * There is  1 square of side 3.
 * Total number of squares = 10 + 4 + 1 = 15.
 *
 *
 * Example 2:
 *
 *
 * Input: matrix =
 * [
 * ⁠ [1,0,1],
 * ⁠ [1,1,0],
 * ⁠ [1,1,0]
 * ]
 * Output: 7
 * Explanation:
 * There are 6 squares of side 1.
 * There is 1 square of side 2.
 * Total number of squares = 6 + 1 = 7.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr.length <= 300
 * 1 <= arr[0].length <= 300
 * 0 <= arr[i][j] <= 1
 *
 *
 */

// @lc code=start
func countSquares(matrix [][]int) int {
	// dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix[0]) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				ans = ans + matrix[i][j]
				continue
			}
			if matrix[i][j] == 1 {
				matrix[i][j] = min(min(matrix[i][j-1], matrix[i-1][j]), matrix[i-1][j-1]) + 1
			} else {
				matrix[i][j] = 0
			}
			ans = ans + matrix[i][j]
		}
	}
	return ans
}

// func min(i, j int) int {
// 	if i > j {
// 		return j
// 	}
// 	return i
// }

// @lc code=end
