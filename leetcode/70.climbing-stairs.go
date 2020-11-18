package leetcode

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	c1, c2 := 1, 2
	for i := 3; i <= n; i++ {
		c1, c2 = c2, c1+c2
	}
	return c2
}

// @lc code=end
