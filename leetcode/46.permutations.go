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
	var backTrack func([]int, int)
	backTrack = func(nums []int, index int) {
		if index == len(nums) {
			resultPermute = append(resultPermute, append([]int{}, nums...))
			return
		}
		for i := index; i < len(nums); i++ {
			// swap
			nums[i], nums[index] = nums[index], nums[i]
			backTrack(nums, index+1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	backTrack(a, 0)
	return resultPermute
}

// @lc code=end
