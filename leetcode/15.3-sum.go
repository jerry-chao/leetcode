package leetcode

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

import "sort"

// @lc code=start
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n == 0 {
		return result
	}
	// sort the input
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[j]+nums[k]+nums[i] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				// skip j = j +1
				for j < len(nums)-1 && j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
				// skip k = k -1
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			} else if nums[j]+nums[k]+nums[i] > 0 {
				// skip k = k -1
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			} else {
				for j < len(nums)-1 && j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			}
		}
		// skip equal i
		for i < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return result
}

// @lc code=end
