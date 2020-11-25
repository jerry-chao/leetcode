package leetcode

/*
 * @lc app=leetcode id=980 lang=golang
 *
 * [980] Unique Paths III
 *
 * https://leetcode.com/problems/unique-paths-iii/description/
 *
 * algorithms
 * Hard (76.94%)
 * Likes:    1139
 * Dislikes: 83
 * Total Accepted:    59.1K
 * Total Submissions: 76.8K
 * Testcase Example:  '[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]'
 *
 * On a 2-dimensional grid, there are 4 types of squares:
 *
 *
 * 1 represents the starting square.  There is exactly one starting square.
 * 2 represents the ending square.  There is exactly one ending square.
 * 0 represents empty squares we can walk over.
 * -1 represents obstacles that we cannot walk over.
 *
 *
 * Return the number of 4-directional walks from the starting square to the
 * ending square, that walk over every non-obstacle square exactly once.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
 * Output: 2
 * Explanation: We have the following two paths:
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
 * 2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
 * Output: 4
 * Explanation: We have the following four paths:
 * 1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
 * 2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
 * 3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
 * 4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
 *
 *
 * Example 3:
 *
 *
 * Input: [[0,1],[2,0]]
 * Output: 0
 * Explanation:
 * There is no path that walks over every empty square exactly once.
 * Note that the starting and ending square can be anywhere in the
 * grid.
 *
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= grid.length * grid[0].length <= 20
 *
 */

// @lc code=start
const (
	startSquare    = 1
	endSquare      = 2
	emptySquare    = 0
	obstacleSquare = -1
)

func uniquePathsIII(grid [][]int) int {
	start := [2]int{}
	emptySquares := 1 // include startSquare
	for v := 0; v < len(grid); v++ {
		for h := 0; h < len(grid[0]); h++ {
			if grid[v][h] == startSquare {
				start = [2]int{v, h}
			}
			if grid[v][h] == emptySquare {
				emptySquares++
			}
		}
	}
	var dfs func(int, int, int, int) int
	dfs = func(v, h int, length int, out int) int {
		if v < 0 || v >= len(grid) || h < 0 || h >= len(grid[0]) {
			return out
		}
		if grid[v][h] == obstacleSquare {
			return out
		}
		if grid[v][h] == endSquare {
			if length == emptySquares {
				return out + 1
			}
			return out
		}
		backtrackSquare := grid[v][h]
		grid[v][h] = obstacleSquare
		out = dfs(v-1, h, length+1, out)
		out = dfs(v+1, h, length+1, out)
		out = dfs(v, h-1, length+1, out)
		out = dfs(v, h+1, length+1, out)
		grid[v][h] = backtrackSquare
		return out
	}
	return dfs(start[0], start[1], 0, 0)
}

// @lc code=end
