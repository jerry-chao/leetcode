package leetcode

/*
 * @lc app=leetcode id=1021 lang=golang
 *
 * [1021] Remove Outermost Parentheses
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	result := ""
	leftNum := 0
	for _, str := range S {
		if string(str) == "(" {
			if leftNum != 0 {
				result = result + "("
			}
			leftNum++
		} else {
			if leftNum > 1 {
				result = result + ")"
			}
			leftNum--
		}
	}
	return result
}

// @lc code=end
