package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

// @lc code=end

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	fmt.Println(reverse(1534236469))
	fmt.Println(-123 % 10)
}
