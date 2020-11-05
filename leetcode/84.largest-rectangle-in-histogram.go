package leetcode

import "log"

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	maxArea := 0
	if len(heights) == 0 {
		return maxArea
	}
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	stack := []int{}
	for i := 0; i < n; i++ {
		// last element
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		log.Println(i, right[i], left[i])
		maxArea = max(maxArea, (right[i]-left[i]-1)*heights[i])
	}
	return maxArea
}

// func max(i, j int) int {
// 	if i > j {
// 		return i
// 	}
// 	return j
// }

// @lc code=end
