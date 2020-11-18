package leetcode

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	if n < 1 {
		return []string{}
	}
	result = []string{}
	generateParenthesisLocal(n, 0, 0, "")
	return result
}

var result []string

func generateParenthesisLocal(n, left, right int, str string) {
	if left == n && right == n {
		result = append(result, str)
		return
	}
	if left < n {
		generateParenthesisLocal(n, left+1, right, str+"(")
	}
	if right < left {
		generateParenthesisLocal(n, left, right+1, str+")")
	}
}

// @lc code=end
