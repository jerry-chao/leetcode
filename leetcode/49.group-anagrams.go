package leetcode

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (57.97%)
 * Likes:    4319
 * Dislikes: 207
 * Total Accepted:    781.7K
 * Total Submissions: 1.3M
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings strs, group the anagrams together. You can return
 * the answer in any order.
 *
 * An Anagram is a word or phrase formed by rearranging the letters of a
 * different word or phrase, typically using all the original letters exactly
 * once.
 *
 *
 * Example 1:
 * Input: strs = ["eat","tea","tan","ate","nat","bat"]
 * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
 * Example 2:
 * Input: strs = [""]
 * Output: [[""]]
 * Example 3:
 * Input: strs = ["a"]
 * Output: [["a"]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 10^4
 * 0 <= strs[i].length <= 100
 * strs[i] consists of lower-case English letters.
 *
 *
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	hash := map[string][]string{}
	for _, str := range strs {
		found := false
		for hashKey, hashValue := range hash {
			if isAnagram(str, hashKey) {
				hash[hashKey] = append(hashValue, str)
				found = true
				break
			}
		}
		if found == false {
			hash[str] = []string{str}
		}
	}
	result := [][]string{}
	for _, strList := range hash {
		result = append(result, strList)
	}
	return result
}

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
