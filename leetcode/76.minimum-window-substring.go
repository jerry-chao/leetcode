package leetcode

import "math"

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (35.20%)
 * Likes:    5463
 * Dislikes: 374
 * Total Accepted:    457.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given two strings s and t, return the minimum window in s which will contain
 * all the characters in t. If there is no such window in s that covers all
 * characters in t, return the empty string "".
 *
 * Note that If there is such a window, it is guaranteed that there will always
 * be only one unique minimum window in s.
 *
 *
 * Example 1:
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Example 2:
 * Input: s = "a", t = "a"
 * Output: "a"
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length, t.length <= 10^5
 * s and t consist of English letters.
 *
 *
 *
 * Follow up: Could you find an algorithm that runs in O(n) time?
 */

// @lc code=start
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// @lc code=end
