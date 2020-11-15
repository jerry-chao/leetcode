package leetcode

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (35.73%)
 * Likes:    3187
 * Dislikes: 2003
 * Total Accepted:    854.9K
 * Total Submissions: 2.4M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] consists of only lower-case English letters.
 *
 *
 */

// @lc code=start

// LongestCommonPrefix with longest common prefix
func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result := ""
	for i := 0; i < len(strs[0]); i++ {
		pre := string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 || string(strs[j][i]) != pre {
				return result
			}
		}
		result = result + pre
	}
	return result
}

// @lc code=end
