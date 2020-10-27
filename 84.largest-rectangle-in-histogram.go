package main

import "fmt"

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	num := len(heights)
	if num == 0 {
		return 0
	}
	left, right := make([]int, num), make([]int, num)
	for i := 0; i < num; i++ {
		right[i] = num
	}
	stack := []int{}
	for i := 0; i < num; i++ {
		// pop last elment
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if heights[i] <= heights[top] {
				stack = stack[:len(stack)-1]
				right[top] = i
			} else {
				break
			}
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	maxArea := 0
	for i := 0; i < num; i++ {
		maxArea = max(maxArea, heights[i]*(right[i]-left[i]-1))
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
	// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{0, 9}))
	fmt.Println(largestRectangleArea([]int{5, 4, 1, 2}))
	fmt.Println(largestRectangleArea([]int{1, 1}))
	fmt.Println(largestRectangleArea([]int{4, 2, 0, 3, 2, 5}))
}
