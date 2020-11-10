package leetcode

import (
	"log"
	"strconv"
)

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
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	if nc == 0 {
		return 0
	}
	result := 0
	visited = make(map[string]int)
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			// if grid is land
			if grid[i][j] == '1' {
				result++
				// remove island to 0
				grid[i][j] = 0
				dfsIslands(grid, nr, nc, i, j)
			}
		}
	}
	return result
}

type direction struct {
	dr int
	dc int
}

var visited map[string]int

func dfsIslands(grid [][]byte, m, n, r, c int) {
	direction := []direction{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < len(direction); i++ {
		dr := direction[i].dr + r
		dc := direction[i].dc + c
		if dr >= 0 && dr < m && dc >= 0 && dc < n && grid[dr][dc] == '1' {
			if _, ok := visited[strconv.Itoa(dr)+strconv.Itoa(dc)]; !ok {
				grid[dr][dc] = '0'
				log.Println(dr, dc)
				dfsIslands(grid, m, n, dr, dc)
			}
		}
	}
}

// @lc code=end
