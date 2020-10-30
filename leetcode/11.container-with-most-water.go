package leetcode

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func MaxArea(height []int) int {
	return maxArea(height)
}

func maxArea(height []int) int {

	maxArea := 0
	if len(height) < 2 {
		return maxArea
	}
	i, j := 0, len(height)-1
	for i < j {
		if height[i] > height[j] {
			maxArea = max(maxArea, height[j]*(j-i))
			j--
		} else {
			maxArea = max(maxArea, height[i]*(j-i))
			i++
		}
	}
	return maxArea
}

// @lc code=end
