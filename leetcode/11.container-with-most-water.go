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
	if len(height) < 2 {
		return 0
	}
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] > height[right] {
			maxArea = max(maxArea, (right-left)*height[right])
			right--
		} else {
			maxArea = max(maxArea, (right-left)*height[left])
			left++
		}
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
