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
<<<<<<< HEAD
	visGenerateParenthesis = map[string]bool{}
	resultGenerateParenthesis = []string{}
	dfsGenerateParenthesis("", 0, 0, 3)
	return resultGenerateParenthesis
}

// dfs generate
var visGenerateParenthesis map[string]bool
var resultGenerateParenthesis []string

func dfsGenerateParenthesis(str string, left, right, max int) {
	// terminator
	if left == max && right == max {
		resultGenerateParenthesis = append(resultGenerateParenthesis, str)
		return
	}

	if left < max {
		dfsGenerateParenthesis(str+"(", left+1, right, max)
	}
	if right < left {
		dfsGenerateParenthesis(str+")", left, right+1, max)
=======
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
>>>>>>> modify leetcode to project
	}
}

// @lc code=end
