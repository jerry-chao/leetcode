package leetcode

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
