package leetcode

import "sort"

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	hash := map[string]int{}
// 	for i := 0; i < len(s); i++ {
// 		hash[string(s[i])]++
// 	}
// 	for i := 0; i < len(t); i++ {
// 		if had, ok := hash[string(t[i])]; ok {
// 			if had > 0 {
// 				hash[string(t[i])]--
// 			} else {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

func isAnagram(s string, t string) bool {
	arrS, arrT := []byte(s), []byte(t)
	sort.Slice(arrS, func(i, j int) bool { return arrS[i] < arrS[j] })
	sort.Slice(arrT, func(i, j int) bool { return arrT[i] < arrT[j] })
	return string(arrS) == string(arrT)
}

// @lc code=end
