package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	result := [][]int{}
	for i := 0; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		if len(result) == 0 || start > result[len(result)-1][1] {
			result = append(result, intervals[i])
		} else {
			result[len(result)-1][1] = max(result[len(result)-1][1], end)
		}
	}
	return result
}

// @lc code=end
