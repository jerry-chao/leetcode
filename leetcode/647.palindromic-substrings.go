package leetcode

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 *
 * https://leetcode.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (61.26%)
 * Likes:    3283
 * Dislikes: 129
 * Total Accepted:    223.1K
 * Total Submissions: 363.4K
 * Testcase Example:  '"abc"'
 *
 * Given a string, your task is to count how many palindromic substrings in
 * this string.
 *
 * The substrings with different start indexes or end indexes are counted as
 * different substrings even they consist of same characters.
 *
 * Example 1:
 *
 *
 * Input: "abc"
 * Output: 3
 * Explanation: Three palindromic strings: "a", "b", "c".
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "aaa"
 * Output: 6
 * Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
 *
 *
 *
 *
 * Note:
 *
 *
 * The input string length won't exceed 1000.
 *
 *
 *
 */

// @lc code=start
func countSubstrings(s string) int {
	// dp[n] = dp[n-1] + num(i, n) i in 0, n
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		n := 0
		for j := 0; j <= i; j++ {
			if isPal(s[j : i+1]) {
				n++
			}
		}
		dp[i] = dp[i-1] + n
	}
	return dp[len(s)-1]
}

func isPal(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end
