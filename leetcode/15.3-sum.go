package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); {
		for j, k := i+1, len(nums)-1; j < k; {

			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
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
