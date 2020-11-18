package leetcode

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (35.08%)
 * Likes:    6233
 * Dislikes: 546
 * Total Accepted:    843.6K
 * Total Submissions: 2.4M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * You are given an integer array nums sorted in ascending order, and an
 * integer target.
 *
 * Suppose that nums is rotated at some pivot unknown to you beforehand (i.e.,
 * [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
 *
 * If target is found in the array return its index, otherwise, return -1.
 *
 *
 * Example 1:
 * Input: nums = [4,5,6,7,0,1,2], target = 0
 * Output: 4
 * Example 2:
 * Input: nums = [4,5,6,7,0,1,2], target = 3
 * Output: -1
 * Example 3:
 * Input: nums = [1], target = 0
 * Output: -1
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 5000
 * -10^4 <= nums[i] <= 10^4
 * All values of nums are unique.
 * nums is guranteed to be rotated at some pivot.
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		curr := left + (right-left)>>1
		if nums[curr] == target {
			return curr
		}
		if nums[left] <= nums[curr] {
			if nums[left] <= target && nums[curr] > target {
				right = curr - 1
			} else {
				left = curr + 1
			}
		} else {
			if nums[curr] < target && nums[right] >= target {
				left = curr + 1
			} else {
				right = curr - 1
			}
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

// @lc code=end
