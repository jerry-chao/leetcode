package leetcode

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	result := []int{}
	if len(nums) < 2 {
		return result
	}
	hash := map[int]int{}
	for index, num := range nums {
		had, ok := hash[num]
		if ok {
			return []int{had, index}
		}
		hash[target-num] = index
	}
	return result
}

// @lc code=end
