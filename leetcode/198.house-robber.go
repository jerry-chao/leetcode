package leetcode

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	// dp[i][0] = max(dp[i-1][1], dp[i-1][0])
	// dp[i][1] = dp[i-1][0]

	// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp0 := nums[0]
	dp1 := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp0, dp1 = dp1, max(dp1, dp0+nums[i])
	}
	return dp1
}

// func max(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }

// @lc code=end
