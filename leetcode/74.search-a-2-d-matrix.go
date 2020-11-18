/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

package leetcode

// @lc code=start

// SearchMatrix search matrix value
func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rn := len(matrix)
	cn := len(matrix[0])

	l, r := 0, rn*cn-1
	for l <= r {
		mid := l + (r-l)/2
		if matrix[mid/cn][mid%cn] == target {
			return true
		}
		if matrix[mid/cn][mid%cn] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l >= rn*cn || matrix[l/cn][l%cn] != target {
		return false
	}
	return true
}

// @lc code=end
