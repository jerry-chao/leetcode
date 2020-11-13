package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 *
 * https://leetcode.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (47.40%)
 * Likes:    2471
 * Dislikes: 69
 * Total Accepted:    403K
 * Total Submissions: 835.4K
 * Testcase Example:  '[1,1,2]'
 *
 * Given a collection of numbers, nums, that might contain duplicates, return
 * all possible unique permutations in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,2]
 * Output:
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 8
 * -10 <= nums[i] <= 10
 *
 *
 */

// @lc code=start
var resultPermuteUnique [][]int
var vis []bool

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	resultPermuteUnique = [][]int{}
	vis = make([]bool, len(nums))
	permuteUniqueBacktrace(nums, []int{}, 0)
	return resultPermuteUnique
}

func permuteUniqueBacktrace(nums []int, perm []int, level int) {
	// terminator
	if level == len(nums) {
		resultPermuteUnique = append(resultPermuteUnique, append([]int{}, perm...))
	}
	// drill down
	for i := 0; i < len(nums); i++ {
		if vis[i] || (i > 0 && vis[i-1] == false && nums[i] == nums[i-1]) {
			continue
		}
		perm = append(perm, nums[i])
		vis[i] = true
		permuteUniqueBacktrace(nums, perm, level+1)
		vis[i] = false
		perm = perm[:len(perm)-1]
	}
}

// @lc code=end
