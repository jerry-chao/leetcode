package leetcode

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	// first index with zero
	first := 0
	for first < len(nums) && nums[first] != 0 {
		first++
	}
	second := first
	for second < len(nums) {
		if nums[second] != 0 {
			nums[first] = nums[second]
			first++
		}
		second++
	}
	// reset first + to zero
	for first < len(nums) {
		nums[first] = 0
		first++
	}
}

// @lc code=end
