package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i < n+1; i++ {
		msg := ""
		if i%3 == 0 {
			msg = msg + "Fizz"
		}
		if i%5 == 0 {
			msg = msg + "Buzz"
		}
		if msg == "" {
			msg = strconv.Itoa(i)
		}
		result = append(result, msg)
	}
	return result
}

// @lc code=end
