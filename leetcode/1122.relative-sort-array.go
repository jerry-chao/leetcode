package leetcode

import "sort"

/*
 * @lc app=leetcode id=1122 lang=golang
 *
 * [1122] Relative Sort Array
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for i := 0; i < len(arr2); i++ {
		rank[arr2[i]] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, okX := rank[x]
		rankY, okY := rank[y]
		if okX && okY {
			return rankX < rankY
		}
		if okX || okY {
			return okX
		}
		return x < y
	})
	return arr1
}

// @lc code=end
