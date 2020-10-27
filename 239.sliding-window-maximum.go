package main

import "fmt"

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
//

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return []int{}
	}

	window, res := []int{}, []int{}
	for i, x := range nums {

		if i >= k && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] <= x {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		fmt.Println(i, x, window)
		if i >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}

// @lc code=end
func main() {
	// nums := []int{1}
	nums := []int{7, 2, 4}
	for _, num := range maxSlidingWindow(nums, 2) {
		fmt.Print(num)
		fmt.Print(" ")
	}
}
