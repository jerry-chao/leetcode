package leetcode

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	hash := map[int]int{}
	for _, num1 := range nums1 {
		hash[num1]++
	}

	for _, num2 := range nums2 {
		had, ok := hash[num2]
		if ok && had > 0 {
			result = append(result, num2)
			hash[num2]--
		}
	}
	return result
}

// @lc code=end
