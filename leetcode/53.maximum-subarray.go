package leetcode

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (47.08%)
 * Likes:    9639
 * Dislikes: 465
 * Total Accepted:    1.2M
 * Total Submissions: 2.5M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Follow up: If you have figured out the O(n) solution, try coding another
 * solution using the divide and conquer approach, which is more subtle.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1]
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [0]
 * Output: 0
 *
 *
 * Example 4:
 *
 *
 * Input: nums = [-1]
 * Output: -1
 *
 *
 * Example 5:
 *
 *
 * Input: nums = [-2147483647]
 * Output: -2147483647
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// dp [i] = max (dp[i-1]+nums[i], nums[i])
	pre, result := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		result = max(pre, result)
	}
	return result
}

// func max(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }

// @lc code=end
