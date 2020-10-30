package leetcode

import "strconv"

/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start

type FizzMap struct {
	Val int
	Str string
}

var hash []FizzMap = []FizzMap{
	FizzMap{Val: 3, Str: "Fizz"},
	FizzMap{Val: 5, Str: "Buzz"},
}

func fizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		result = append(result, fizzBuzzString(i))
	}
	return result
}

func fizzBuzzString(n int) string {
	result := ""
	for _, value := range hash {
		if n%value.Val == 0 {
			result = result + value.Str
		}
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result
}

// @lc code=end
