package leetcode

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
var ans [][]int

func permute(a []int) [][]int {
	ans = make([][]int, 0)
	perm(a, 0)
	return ans
}

func perm(a []int, i int) {
	if i == len(a) {
		ans = append(ans, append(make([]int, 0), a...))
		return
	}
	for j := i; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// @lc code=end
