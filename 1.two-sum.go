/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package main

/* func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
} */

// @lc code=start
func twoSum(nums []int, target int) []int {
	result := []int{}
	if len(nums) < 2 {
		return result
	}
	hash := make(map[int]int)
	for index, num := range nums {
		numIndex, ok := hash[num]
		if ok {
			return []int{numIndex, index}
		}
		hash[target-num] = index
	}
	return result
}

// @lc code=end
