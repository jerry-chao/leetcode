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
	}
}

// @lc code=end
