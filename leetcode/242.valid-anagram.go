package leetcode

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := map[string]int{}
	for i := 0; i < len(s); i++ {
		hash[string(s[i])]++
	}
	for i := 0; i < len(t); i++ {
		if had, ok := hash[string(t[i])]; ok {
			if had > 0 {
				hash[string(t[i])]--
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// @lc code=end
