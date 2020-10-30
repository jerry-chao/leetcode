package leetcode

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}
// 	first, second := 1, 2
// 	var tmp int
// 	for i := 3; i <= n; i++ {
// 		tmp = second
// 		second = first + tmp
// 		first = tmp
// 	}
// 	return second
// }

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	first, second := 1, 2
	for i := 2; i < n; i++ {
		tmp := second
		second = first + second
		first = tmp
	}
	return second
}

// @lc code=end
