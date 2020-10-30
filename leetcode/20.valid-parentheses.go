package main

import "fmt"

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := []string{}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if isLeft(char) {
			stack = append(stack, transfer(char))
			continue
		}

		if len(stack) > 0 && char == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func isLeft(char string) bool {
	if char == "{" || char == "[" || char == "(" {
		return true
	}
	return false
}

func transfer(char string) string {
	if char == "{" {
		return "}"
	} else if char == "(" {
		return ")"
	} else {
		return "]"
	}
}

// @lc code=end

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("["))
}
