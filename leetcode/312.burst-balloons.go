package leetcode

/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 *
 * https://leetcode.com/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (52.74%)
 * Likes:    2888
 * Dislikes: 74
 * Total Accepted:    109.9K
 * Total Submissions: 207.6K
 * Testcase Example:  '[3,1,5,8]'
 *
 * Given n balloons, indexed from 0 to n-1. Each balloon is painted with a
 * number on it represented by array nums. You are asked to burst all the
 * balloons. If the you burst balloon i you will get nums[left] * nums[i] *
 * nums[right] coins. Here left and right are adjacent indices of i. After the
 * burst, the left and right then becomes adjacent.
 *
 * Find the maximum coins you can collect by bursting the balloons wisely.
 *
 * Note:
 *
 *
 * You may imagine nums[-1] = nums[n] = 1. They are not real therefore you can
 * not burst them.
 * 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
 *
 *
 * Example:
 *
 *
 * Input: [3,1,5,8]
 * Output: 167
 * Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  -->
 * []
 * coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
 *
 */

// @lc code=start
func maxCoins(nums []int) int {
	n := len(nums)
	rec := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		rec[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += rec[i][k] + rec[k][j]
				rec[i][j] = max(rec[i][j], sum)
			}
		}
	}
	return rec[0][n+1]
}

// func max(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }

// @lc code=end
