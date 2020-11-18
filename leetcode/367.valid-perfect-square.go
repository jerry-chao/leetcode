package leetcode

import "math"

/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 *
 * https://leetcode.com/problems/valid-perfect-square/description/
 *
 * algorithms
 * Easy (41.87%)
 * Likes:    1030
 * Dislikes: 178
 * Total Accepted:    243.7K
 * Total Submissions: 581.5K
 * Testcase Example:  '16'
 *
 * Given a positive integer num, write a function which returns True if num is
 * a perfect square else False.
 *
 * Follow up: Do not use any built-in library function such as sqrt.
 *
 *
 * Example 1:
 * Input: num = 16
 * Output: true
 * Example 2:
 * Input: num = 14
 * Output: false
 *
 *
 * Constraints:
 *
 *
 * 1 <= num <= 2^31 - 1
 *
 *
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	sqrt := sqrtSquare(num)
	if sqrt*sqrt == num {
		return true
	}
	return false
}

func sqrtSquare(num int) int {
	curr := float64(1)
	target := float64(num)
	for {
		pre := curr
		curr = (curr + target/curr) * 0.5
		if math.Abs(curr-pre) < 1e-7 {
			break
		}
	}
	return int(curr)
}

// @lc code=end
