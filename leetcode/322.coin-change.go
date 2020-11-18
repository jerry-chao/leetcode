package leetcode

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (36.26%)
 * Likes:    5368
 * Dislikes: 163
 * Total Accepted:    505.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 *
 * You may assume that you have an infinite number of each kind of coin.
 *
 *
 * Example 1:
 *
 *
 * Input: coins = [1,2,5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Example 3:
 *
 *
 * Input: coins = [1], amount = 0
 * Output: 0
 *
 *
 * Example 4:
 *
 *
 * Input: coins = [1], amount = 1
 * Output: 1
 *
 *
 * Example 5:
 *
 *
 * Input: coins = [1], amount = 2
 * Output: 2
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= coins.length <= 12
 * 1 <= coins[i] <= 2^31 - 1
 * 0 <= amount <= 10^4
 *
 *
 */

// @lc code=start
// dp[n] = min(dp[n-k]) + 1 k in coins
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dpCoinChange = make([]int, amount+1)
	dpCoinChange[0] = 0

	for i := 1; i <= amount; i++ {
		minChange := math.MaxInt64
		for _, coin := range coins {
			if i >= coin {
				minChange = min(minChange, dpCoinChange[i-coin])
			}
		}
		if minChange != math.MaxInt64 {
			dpCoinChange[i] = minChange + 1
		} else {
			dpCoinChange[i] = math.MaxInt64
		}
	}
	if dpCoinChange[amount] == math.MaxInt64 {
		return -1
	}
	return dpCoinChange[amount]
}

// func min(i, j int) int {
// 	if i > j {
// 		return j
// 	}
// 	return i
// }

var dpCoinChange []int

// @lc code=end
