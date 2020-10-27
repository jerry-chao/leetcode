package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	first := 0
	second := len(height) - 1
	maxArea := 0
	for first < second {
		if height[first] > height[second] {
			maxArea = max(maxArea, (second-first)*height[second])
			second--
		} else {
			maxArea = max(maxArea, (second-first)*height[first])
			first++
		}
	}
	return maxArea
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
