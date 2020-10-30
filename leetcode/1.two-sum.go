package main

import "fmt"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

/* func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
} */

// @lc code=start
func twoSum(nums []int, target int) []int {
	result := []int{}
	if len(nums) < 2 {
		return
	}
	hash := map[int]int{}
	for _, num := range nums {
		had, ok := hash[num]
		if ok {

		}
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
