package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println(lengthOfLongestSubstring("test"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

//  calc length of longest substring
func lengthOfLongestSubstring(s string) int {
	charlist := make(map[rune]int)
	maxlen, index, tmp := 0, 0, 0
	for idx, value := range s {
		vidx, ok := charlist[value]
		if ok && vidx >= index {
			tmp = idx - index
			maxlen = max(maxlen, tmp)
			index = vidx + 1
		}
		charlist[value] = idx
	}
	fmt.Println(maxlen, index)
	return max(maxlen, len(s)-index)
}

// max ...
func max(first, second int) int {
	if first < second {
		return second
	}
	return first
}
