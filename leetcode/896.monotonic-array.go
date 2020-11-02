package leetcode

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
	status := 0
	for i := 1; i < len(A); i++ {
		status1 := transfer(A[i-1], A[i])
		if status == 0 {
			status = status1
		} else if status1 == 0 || status1 == status {
			continue
		} else {
			return false
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
