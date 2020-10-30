package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return reversePalindrome(x) == x
}

func reversePalindrome(x int) int {
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
