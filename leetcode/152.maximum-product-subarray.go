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
	minValue, maxValue, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax, tmpMin := maxValue, minValue
		maxValue = max(max(tmpMax*nums[i], tmpMin*nums[i]), nums[i])
		minValue = min(min(tmpMax*nums[i], tmpMin*nums[i]), nums[i])
		ans = max(maxValue, ans)
	}

	return ans
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
