package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (25.28%)
 * Likes:    3445
 * Dislikes: 3225
 * Total Accepted:    469.6K
 * Total Submissions: 1.8M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z is being encoded to numbers using the
 * following mapping:
 *
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * Given a non-empty string containing only digits, determine the total number
 * of ways to decode it.
 *
 * The answer is guaranteed to fit in a 32-bit integer.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "12"
 * Output: 2
 * Explanation: It could be decoded as "AB" (1 2) or "L" (12).
 *
 *
 * Example 2:
 *
 *
 * Input: s = "226"
 * Output: 3
 * Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2
 * 6).
 *
 *
 * Example 3:
 *
 *
 * Input: s = "0"
 * Output: 0
 * Explanation: There is no character that is mapped to a number starting with
 * '0'. We cannot ignore a zero when we face it while decoding. So, each '0'
 * should be part of "10" --> 'J' or "20" --> 'T'.
 *
 *
 * Example 4:
 *
 *
 * Input: s = "1"
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 100
 * s contains only digits and may contain leading zero(s).
 *
 *
 */

// @lc code=start
func numDecodings(s string) int {
	// dp[i] = dp[i-1] + dp[i-2] (dp[i-1]dp[i] is valid)
	dp := make([]int, len(s)+1)
	code, _ := strconv.Atoi(string(s[0]))
	if code <= 0 {
		return 0
	}
	dp[0], dp[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		currentCode, _ := strconv.Atoi(string(s[i-1 : i]))
		allCode, _ := strconv.Atoi(string(s[i-2 : i]))
		if currentCode == 0 && (allCode == 0 || allCode > 20) {
			return 0
		} else if currentCode == 0 {
			dp[i] = dp[i-2]
		} else if allCode < 10 {
			dp[i] = dp[i-1]
		} else if allCode > 26 {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[len(s)]
}

// @lc code=end
