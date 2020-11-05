package leetcode

import "log"

/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */
//

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	result := []int{}
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			log.Println("pop index", i)
			// pop elment less than nums[i]
			stack = stack[:len(stack)-1]
		}
		// if len(stack) > k
		if len(stack) > 0 && stack[0]+k-1 < i {
			stack = stack[1:]
		}
		stack = append(stack, i)
		if i >= k-1 {
			result = append(result, nums[stack[0]])
		}
	}
	return result
}

// @lc code=end
