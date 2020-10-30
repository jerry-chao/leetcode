package main

import "fmt"

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
	for _, schar := range s {
		hash[string(schar)]++
	}
	for _, tchar := range t {
		had, ok := hash[string(tchar)]
		if ok && had > 0 {
			hash[string(tchar)]--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}
