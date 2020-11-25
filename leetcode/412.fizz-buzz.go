package leetcode

import "strconv"

/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start

func fizzBuzz(n int) []string {
	return fizzBuzzRec(1, n, []string{})
}
func fizzBuzzRec(n, max int, result []string) []string {
	// terminator
	if n > max {
		return result
	}
	// current process
	result = append(result, int2Str(n))

	// drill down
	return fizzBuzzRec(n+1, max, result)
}

// Fizz val map to fizz string
type Fizz struct {
	Val int
	Str string
}

func int2Str(n int) string {
	// map is not ordered
	hash := []Fizz{
		{Val: 3, Str: "Fizz"},
		{Val: 5, Str: "Buzz"},
	}
	result := ""
	for _, fizz := range hash {
		if n%fizz.Val == 0 {
			result = result + fizz.Str
		}
	}
	if result == "" {
		return strconv.Itoa(n)
	}
	return result
}

// @lc code=end
