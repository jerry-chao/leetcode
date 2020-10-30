package main

import "leetcode/leetcode"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	leetcode.Rotate(nums, k)
	leetcode.PrintArray(nums)

	nums = []int{1, 2, 3, 4, 5, 6}
	k = 4
	leetcode.Rotate(nums, k)
	leetcode.PrintArray(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	leetcode.Rotate(nums, k)
	leetcode.PrintArray(nums)

	nums = []int{1, 2}
	k = 2
	leetcode.Rotate(nums, k)
	leetcode.PrintArray(nums)
}
