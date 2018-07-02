package main

import (
	"fmt"
)

// main ...
func main() {
	fmt.Println("hello world")
	fmt.Println("max1", maxArea([]int{1, 2, 4, 3}))
	fmt.Println("max1", maxArea([]int{1, 1}))
	fmt.Println("max1", maxArea([]int{1, 2, 1}))
	fmt.Println("max1", maxArea([]int{1, 2, 4}))
	fmt.Println("max1", maxArea([]int{4, 5, 6}))
}

// maxArea ...
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	result := 0
	start := 0
	end := len(height) - 1
	for start < end {
		result = max(min(height[start], height[end])*(end-start), result)
		fmt.Println("this is result", result)
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return result
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}
