package leetcode

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}
	result := []int{}
	hash := map[int]int{}
	for _, num1 := range nums1 {
		hash[num1] = 1
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
