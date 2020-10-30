package leetcode

func AddDigits(num int) int {
	return addDigits(num)
}

/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
func addDigits(num int) int {
	// return addDigitsO1(num)
	return addDigitsLoops(num)
}

func addDigitsLoops(num int) int {
	if num < 10 {
		return num
	}
	target := 0

	for num >= 10 {
		target = target + num%10
		num = num / 10
	}

	return addDigitsLoops(target + num)
}

func addDigitsO1(num int) int {
	return (num-1)%9 + 1
}

// @lc code=end
