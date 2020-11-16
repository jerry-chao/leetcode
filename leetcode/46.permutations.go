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
	if len(a) < 1 {
		return resultPermute
	}
	var backTrack func([]int, int)
	backTrack = func(a []int, index int) {
		if index == len(a) {
			resultPermute = append(resultPermute, append([]int{}, a...))
			return
		}
		for i := index; i < len(a); i++ {
			a[index], a[i] = a[i], a[index]
			backTrack(a, index+1)
			a[index], a[i] = a[i], a[index]
		}
	}
	backTrack(a, 0)
	return resultPermute
}

// @lc code=end
