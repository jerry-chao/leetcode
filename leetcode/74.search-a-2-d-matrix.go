/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

package leetcode

import "fmt"

// @lc code=start
func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 {
			return false
		}
		if matrix[i][0] > target {
			return false
		} else if matrix[i][len(matrix[i])-1] >= target {
			return searchTarget(matrix[i], target)
		} else {
			continue
		}
	}

	return false
}

func searchTarget(nums []int, target int) bool {
	start := 0
	end := len(nums)
	fmt.Println("targetArrayLen:", len(nums), nums[0])
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

// @lc code=end
