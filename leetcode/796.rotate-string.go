package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode id=796 lang=golang
 *
 * [796] Rotate String
 */

// @lc code=start
func rotateString(A string, B string) bool {
	if len(A) == len(B) {
		return strings.Contains(A+A, B)
	} else {
		return false
	}
}

// @lc code=end
