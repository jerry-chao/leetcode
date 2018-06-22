package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("ab", "a"))
	fmt.Println(isAnagram("中那个", "那中个"))
}

func isAnagram(s string, t string) bool {
	sHash := string2hash(s)
	for _, value := range t {
		sHash[value] = sHash[value] - 1
		if sHash[value] < 0 {
			return false
		} else if sHash[value] == 0 {
			delete(sHash, value)
		}
	}
	return len(sHash) == 0
}

func string2hash(s string) map[rune]int {
	ret := make(map[rune]int)
	for _, value := range s {
		ret[value] = ret[value] + 1
	}
	return ret
}
