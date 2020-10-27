package main

import "fmt"

/*
 * @lc app=leetcode id=896 lang=golang
 *
 * [896] Monotonic Array
 */

// @lc code=start
func isMonotonic(A []int) bool {
	if len(A) < 3 {
		return true
	}
	incr := 0
	for i := 1; i < len(A); i++ {
		tmp := transfer(A[i], A[i-1])
		if tmp == 0 {
			continue
		}
		if incr != tmp && incr != 0 {
			return false
		}
		if incr == 0 {
			incr = tmp
		}
	}
	return true
}

func transfer(i, j int) int {
	if i > j {
		return 1
	} else if i < j {
		return -1
	} else {
		return 0
	}
}

// @lc code=end

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(isMonotonic([]int{1, 3, 2}))
	fmt.Println(isMonotonic([]int{1, 2, 4, 5}))
	fmt.Println(isMonotonic([]int{1, 1, 1}))
}
