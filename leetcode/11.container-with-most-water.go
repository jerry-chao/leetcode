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
