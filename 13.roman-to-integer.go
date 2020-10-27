package main

import "fmt"

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start

func romanToInt(s string) int {
	result := 0
	pre := ""
	for _, str := range s {
		if pre != "" {
			multi, multiPlus := transfer(pre, string(str))
			if multi {
				result = result + multiPlus
				pre = ""
				continue
			} else {
				result = result + transferNum(pre)
				pre = ""
			}
		}
		if isFirst(string(str)) {
			pre = string(str)
		} else {
			result = result + transferNum(string(str))
		}

	}
	if pre != "" {
		result = result + transferNum(pre)

	}
	return result
}

func isFirst(str string) bool {
	if str == "I" || str == "X" || str == "C" {
		return true
	}
	return false
}

func transfer(pre, str string) (bool, int) {
	if pre == "I" && str == "V" {
		return true, 4
	} else if pre == "I" && str == "X" {
		return true, 9
	} else if pre == "X" && str == "L" {
		return true, 40
	} else if pre == "X" && str == "C" {
		return true, 90
	} else if pre == "C" && str == "D" {
		return true, 400
	} else if pre == "C" && str == "M" {
		return true, 900
	} else {
		return false, 0
	}
}

func transferNum(str string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	return roman[str]
}

// @lc code=end

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
}
