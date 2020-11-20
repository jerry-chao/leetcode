package leetcode

import "math"

/*
 * @lc app=leetcode id=363 lang=golang
 *
 * [363] Max Sum of Rectangle No Larger Than K
 *
 * https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/description/
 *
 * algorithms
 * Hard (38.01%)
 * Likes:    934
 * Dislikes: 70
 * Total Accepted:    49.5K
 * Total Submissions: 129.6K
 * Testcase Example:  '[[1,0,1],[0,-2,3]]\n2'
 *
 * Given a non-empty 2D matrix matrix and an integer k, find the max sum of a
 * rectangle in the matrix such that its sum is no larger than k.
 *
 * Example:
 *
 *
 * Input: matrix = [[1,0,1],[0,-2,3]], k = 2
 * Output: 2
 * Explanation: Because the sum of rectangle [[0, 1], [-2, 3]] is
 * 2,
 * and 2 is the max number no larger than k (k = 2).
 *
 * Note:
 *
 *
 * The rectangle inside the matrix must have an area > 0.
 * What if the number of rows is much larger than the number of columns?
 *
 */

// @lc code=start
func maxSumSubmatrix(matrix [][]int, k int) int {
	rowNum, colNum := len(matrix), len(matrix[0])
	//因为行远大于行，所以按行从左到右求前缀和
	for row := 0; row < rowNum; row++ {
		for col := 1; col < colNum; col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}
	result := math.MinInt32
	// 先固定左右边界，移动上下边界，求和
	for left := 0; left < colNum; left++ {
		for right := left; right < colNum; right++ {
			for top := 0; top < rowNum; top++ {
				sum := 0
				for bottom := top; bottom < rowNum; bottom++ {
					if left == 0 {
						sum += matrix[bottom][right]
					} else {
						sum += matrix[bottom][right] - matrix[bottom][left-1]
					}
					if sum <= k && sum > result {
						result = sum
					}
				}
			}
		}
	}
	return result
}

// @lc code=end
