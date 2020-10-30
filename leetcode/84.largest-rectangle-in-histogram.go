package leetcode

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights)
	if n == 0 {
		return maxArea
	}
	// init left and right
	// left index before height
	// right index after height
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	stack := []int{}
	for i := 0; i < n; i++ {
		// pop stack
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
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
	// caculator maxArea
	for i := 0; i < n; i++ {
		maxArea = max(maxArea, heights[i]*(right[i]-left[i]-1))
	}
	return maxArea
}

// @lc code=end