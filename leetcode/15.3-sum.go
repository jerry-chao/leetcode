/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package leetcode

import (
	"fmt"
	"sort"
)

// @lc code=start
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n == 0 {
		return result
	}
	sort.Ints(nums)
	i := 0
	for i < n {
		target := -nums[i]
		j := i + 1
		k := n - 1
		fmt.Println(j, k)
		for j < k {
			if nums[j]+nums[k] == target {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k {
					if nums[j] != nums[j-1] {
						break
					}
					j++
				}
				k--
				for j < k {
					if nums[k] != nums[k+1] {
						break
					}
					k--
				}
			} else if nums[j]+nums[k] > target {
				k--
				for j < k {
					if nums[k] != nums[k+1] {
						break
					}
					k--
				}
			} else {
				j++
				for j < k {
					if nums[j] != nums[j-1] {
						break
					}
					j++
				}
			}
		}
		i++
		for i < n-1 {
			if nums[i] != nums[i-1] {
				break
			}
			i++
		}
	}
	return result
}

// @lc code=end
