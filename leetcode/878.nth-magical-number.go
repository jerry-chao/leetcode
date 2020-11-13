package leetcode

/*
 * @lc app=leetcode id=878 lang=golang
 *
 * [878] Nth Magical Number
 *
 * https://leetcode.com/problems/nth-magical-number/description/
 *
 * algorithms
 * Hard (28.65%)
 * Likes:    230
 * Dislikes: 62
 * Total Accepted:    10.4K
 * Total Submissions: 36.5K
 * Testcase Example:  '1\n2\n3'
 *
 * A positive integer is magical if it is divisible by either A or B.
 *
 * Return the N-th magical number.  Since the answer may be very large, return
 * it modulo 10^9 + 7.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: N = 1, A = 2, B = 3
 * Output: 2
 *
 *
 *
 * Example 2:
 *
 *
 * Input: N = 4, A = 2, B = 3
 * Output: 6
 *
 *
 *
 * Example 3:
 *
 *
 * Input: N = 5, A = 2, B = 4
 * Output: 10
 *
 *
 *
 * Example 4:
 *
 *
 * Input: N = 3, A = 6, B = 4
 * Output: 8
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= N <= 10^9
 * 2 <= A <= 40000
 * 2 <= B <= 40000
 *
 *
 *
 *
 *
 *
 */

// @lc code=start
func nthMagicalNumber(N int, A int, B int) int {
	// BFS
	queue := []int{A, B}
	cnt := 1
	for len(queue) > 0 && cnt < N {
		top := queue[0]
		// pop queue
		queue = queue[1:]
		cnt++
		queue = append(queue, top*A, top*B)
	}
	return queue[0]
}

// @lc code=end
