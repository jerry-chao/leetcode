package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func Reverse(x int) int {
	return reverse(x)
}

func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

// @lc code=end
