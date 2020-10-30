package leetcode

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
// func moveZeroes(nums []int) {
// 	if len(nums) == 0 {
// 		return
// 	}

// 	first := 0
// 	for first < len(nums) {
// 		if nums[first] == 0 {
// 			break
// 		}
// 		first++
// 	}

// 	second := first
// 	for second < len(nums) {
// 		if nums[second] != 0 {
// 			nums[first] = nums[second]
// 			nums[second] = 0
// 			first++
// 		}
// 		second++
// 	}
// }

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	// set first to zero
	first := 0
	for first < len(nums) {
		if nums[first] == 0 {
			break
		}
		first++
	}
	second := first + 1
	for second < len(nums) {
		if nums[second] != 0 {
			nums[first] = nums[second]
			nums[second] = 0
			first++
		}
		second++
	}
}

// @lc code=end
