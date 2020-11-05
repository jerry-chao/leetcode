package leetcode

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
	for _, schar := range s {
		str := transferParentheses(string(schar))
		if str != "" {
			stack = append(stack, str)
		} else {
			if len(stack) > 0 {
				// pop first elment
				if stack[len(stack)-1] == string(schar) {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func transferParentheses(str string) string {
	hash := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	if had, ok := hash[str]; ok {
		return had
	}
	return ""

}

// @lc code=end
