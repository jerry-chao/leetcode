package leetcode

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first := 1
	second := 2
	index := 2
	for index < n {
		tmp := second
		second = first + second
		first = tmp
		index++
	}
	return second
}

// @lc code=end
