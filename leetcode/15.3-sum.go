/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	i := 0
	for i < len(nums) {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
		i++
		for i < len(nums) && nums[i] == nums[i-1] {

			i++
		}
	}
	return result
}

// @lc code=end
