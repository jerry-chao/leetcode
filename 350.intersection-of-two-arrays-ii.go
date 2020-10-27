package main

import "fmt"

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	hash := map[int]int{}

	for _, num1 := range nums1 {
		hash[num1] = hash[num1] + 1
	}
	for _, num2 := range nums2 {
		had, ok := hash[num2]
		if ok && had > 0 {
			hash[num2] = hash[num2] - 1
			result = append(result, num2)
		}
	}
	return result
}

// @lc code=end

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	for _, num := range intersect(nums1, nums2) {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
