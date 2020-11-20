package leetcode

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (32.34%)
 * Likes:    5563
 * Dislikes: 190
 * Total Accepted:    411.2K
 * Total Submissions: 1.3M
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find the contiguous subarray within an array
 * (containing at least one number) which has the largest product.
 *
 * Example 1:
 *
 *
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 *
 *
 * Example 2:
 *
 *
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 *
 */

// @lc code=start
func maxProduct(nums []int) int {
	// max[i] = max(max[i-1]*nums[i], min[i-1]*nums[i], nums[i])
	// min[i] = min(max[i-1]*nums[i], min[i-1]*nums[i], nums[i])
	maxPre, minPre, maxValue := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax := maxPre
		maxPre = max(max(tmpMax*nums[i], minPre*nums[i]), nums[i])
		minPre = min(min(tmpMax*nums[i], minPre*nums[i]), nums[i])
		maxValue = max(maxValue, maxPre)
	}
	return maxValue
}

// func max(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }

// func min(i, j int) int {
// 	if i > j {
// 		return j
// 	}
// 	return i
// }

// @lc code=end
