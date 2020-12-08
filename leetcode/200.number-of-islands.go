package leetcode

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (47.55%)
 * Likes:    6857
 * Dislikes: 220
 * Total Accepted:    864.3K
 * Total Submissions: 1.8M
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given an m x n 2d grid map of '1's (land) and '0's (water), return the
 * number of islands.
 *
 * An island is surrounded by water and is formed by connecting adjacent lands
 * horizontally or vertically. You may assume all four edges of the grid are
 * all surrounded by water.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","1","1","0"],
 * ⁠ ["1","1","0","1","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","0","0","0"]
 * ]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["1","1","0","0","0"],
 * ⁠ ["0","0","1","0","0"],
 * ⁠ ["0","0","0","1","1"]
 * ]
 * Output: 3
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 300
 * grid[i][j] is '0' or '1'.
 *
 *
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	result := 0
	nr := len(grid)
	nc := len(grid[0])
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				result++
				grid[i][j] = '0'
				grid = clearRelations(grid, nr, nc, i, j)
			}
		}
	}
	return result
}

var dRow = []int{0, 0, -1, 1}
var dCol = []int{1, -1, 0, 0}

func clearRelations(grid [][]byte, nr, nc, row, col int) [][]byte {
	for i := 0; i < 4; i++ {
		rowNew := row + dRow[i]
		colNew := col + dCol[i]
		if rowNew < nr && rowNew >= 0 && colNew < nc && colNew >= 0 && grid[rowNew][colNew] == '1' {
			grid[rowNew][colNew] = '0'
			grid = clearRelations(grid, nr, nc, rowNew, colNew)
		}
	}
	return grid
}

// @lc code=end
