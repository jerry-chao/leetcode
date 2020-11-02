package leetcode

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	numsWithFrequent := map[int]int{}
	for i := 0; i < len(nums); i++ {
		numsWithFrequent[nums[i]]++
	}
	if len(numsWithFrequent) < k {
		return []int{}
	}
	return []int{}
}

// @lc code=end
