package main

import "fmt"

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if isFirst(string(s[i])) {
			// push to stack
			stack = append(stack, reverse(string(s[i])))
		} else {
			// pop stack
			if len(stack) == 0 {
				return false
			}
			if string(s[i]) != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isFirst(char string) bool {
	return char == "(" || char == "[" || char == "{"
}

func reverse(char string) string {
	if char == "(" {
		return ")"
	} else if char == "[" {
		return "]"
	} else {
		return "}"
	}
}

// @lc code=end

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("["))
}
