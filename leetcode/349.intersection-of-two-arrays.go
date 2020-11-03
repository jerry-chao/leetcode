package leetcode

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		intersection(nums2, nums1)
	}
	hash1 := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		hash1[nums1[i]]++
	}
	resultMap := map[int]int{}
	for i := 0; i < len(nums2); i++ {
		if _, ok := hash1[nums2[i]]; ok {
			resultMap[nums2[i]] = 1
		}
	}
	result := []int{}
	for key := range resultMap {
		result = append(result, key)
	}
	return result
}

// @lc code=end
