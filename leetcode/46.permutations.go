package leetcode

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
var resultPermute [][]int

func permute(a []int) [][]int {
	resultPermute = [][]int{}
	permuteBacktrace(a, 0)
	return resultPermute
}

func permuteBacktrace(a []int, index int) {
	if len(a) == index {
		// golang share array, so must copy
		resultPermute = append(resultPermute, append([]int{}, a...))
		return
	}
	for i := index; i < len(a); i++ {
		a[i], a[index] = a[index], a[i]
		permuteBacktrace(a, index+1)
		a[i], a[index] = a[index], a[i]
	}
}

// @lc code=end
