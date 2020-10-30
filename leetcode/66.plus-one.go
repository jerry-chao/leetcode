package main

import "fmt"

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	pre := 1
	result := make([]int, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + pre
		pre = sum / 10
		result[i] = sum % 10
	}
	if pre == 1 {
		result = append([]int{pre}, result...)
	}
	return result
}

// @lc code=end

func main() {
	printArray(plusOne([]int{1, 2, 3}))
	printArray(plusOne([]int{4, 3, 2, 1}))
	printArray(plusOne([]int{0}))
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("")
}
