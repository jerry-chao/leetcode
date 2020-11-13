package leetcode

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
var result [][]int

func permute(a []int) [][]int {
	result = [][]int{}
	permuteLocal(a, 0)
	return result
}

func permuteLocal(a []int, index int) {
	if len(a) == index {
		result = append(result, append([]int{}, a...))
	}
	for i := index; i < len(a); i++ {
		a[index], a[i] = a[i], a[index]
		permuteLocal(a, index+1)
		a[i], a[index] = a[index], a[i]
	}
}

// @lc code=end
