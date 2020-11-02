package leetcode

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	num := len(height)
	if num == 0 {
		return 0
	}
	volume := 0
	leftMax := height[0]
	stack := []int{}
	for i := 1; i < len(height); i++ {
		if height[i] >= leftMax {
			// pop stack
			for len(stack) > 0 {
				lastElement := stack[len(stack)-1]
				volume = volume + leftMax - height[lastElement]
				stack = stack[:len(stack)-1]
			}
			leftMax = height[i]
		} else {
			stack = append(stack, i)
		}
	}
	rightMax := 0
	for len(stack) > 0 {
		lastElement := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if height[lastElement] < rightMax {
			volume = volume + rightMax - height[lastElement]
		} else {
			rightMax = height[lastElement]
		}
	}
	return volume

}

// @lc code=end
