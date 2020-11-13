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
	nc := len(grid)
	nr := len(grid[0])
	for c := 0; c < nc; c++ {
		for r := 0; r < nr; r++ {
			if grid[c][r] == '1' {
				result++
				grid[c][r] = '0'
				removeOneFromGrid(grid, c, r, nc, nr)
			}
		}
	}
	return result
}

func removeOneFromGrid(grid [][]byte, c, r, nc, nr int) {
	direction := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dr := range direction {
		cNext := c + dr[0]
		rNext := r + dr[1]
		if cNext >= 0 && rNext >= 0 && cNext < nc && rNext < nr {
			if grid[cNext][rNext] == '1' {
				grid[cNext][rNext] = '0'
				removeOneFromGrid(grid, cNext, rNext, nc, nr)
			}
		}
	}
}

// @lc code=end
