package leetcode

/*
 * @lc app=leetcode id=1021 lang=golang
 *
 * [1021] Remove Outermost Parentheses
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	stack := []string{}
	result := ""
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "(" && len(stack) > 0 {
			result = result + string(S[i])
			stack = append(stack, ")")
		} else if string(S[i]) == "(" {
			stack = append(stack, ")")
		} else if len(stack) > 0 {
			// pop top of stack
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				result = result + string(S[i])
			}
		} else {
			return ""
		}
	}
	return result
}

// @lc code=end
