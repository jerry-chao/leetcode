/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

package leetcode

import "fmt"

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		fmt.Println("found group:", i)
		if len(matrix[i]) > 0 && matrix[i][0] > target {
			if i > 0 {
				return search(matrix[i-1], target)
			} else {
				return false
			}
		}
	}
	return search(matrix[len(matrix)-1], target)
}

func search(nums []int, target int) bool {
	fmt.Println("xxx")
	if len(nums) == 0 {
		return false
	}
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		fmt.Println("mid:", mid, nums[mid])
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return false
}

// @lc code=end
