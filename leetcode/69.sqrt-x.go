package leetcode

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (34.42%)
 * Likes:    1603
 * Dislikes: 2101
 * Total Accepted:    627.4K
 * Total Submissions: 1.8M
 * Testcase Example:  '4'
 *
 * Given a non-negative integer x, compute and return the square root of x.
 *
 * Since the return type is an integer, the decimal digits are truncated, and
 * only the integer part of the result is returned.
 *
 *
 * Example 1:
 *
 *
 * Input: x = 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: x = 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since the decimal part
 * is truncated, 2 is returned.
 *
 *
 * Constraints:
 *
 *
 * 0 <= x <= 2^31 - 1
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// func mySqrt(x int) int {
// 	target := float64(x)
// 	curr := float64(1)
// 	for {
// 		pre := curr
// 		curr = (curr + target/curr) * 0.5
// 		if math.Abs(curr-pre) < 1e-7 {
// 			break
// 		}
// 	}
// 	return int(curr)
// }

// @lc code=end
