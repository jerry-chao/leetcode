package main

import "fmt"

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// two point from last of nums1 and nums2
	i := m - 1
	j := n - 1
	index := len(nums1) - 1
	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[index] = nums2[j]
			j--
		} else {
			nums1[index] = nums1[i]
			i--
		}
		index--
	}
	for i := 0; i <= j; i++ {
		nums1[i] = nums2[i]
	}
}

// @lc code=end

func main() {
	nums1 := []int{7, 8, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 2, nums2, 3)
	printArray(nums1)
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("")
}
