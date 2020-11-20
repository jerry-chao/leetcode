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
	dp0, dp1 := 1, 2
	for i := 3; i <= n; i++ {
		dp0, dp1 = dp1, dp0+dp1
	}
	return dp1
}

// @lc code=end
