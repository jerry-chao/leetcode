package leetcode

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	// double pointer
	if len(nums) < 2 {
		return len(nums)
	}
	first := 0
	for second := first + 1; second < len(nums); second++ {
		if nums[second] == nums[first] {
			continue
		}
		first++
		nums[first] = nums[second]
	}
	return first + 1
}

// @lc code=end
