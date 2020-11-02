package leetcode

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=299 lang=golang
 *
 * [299] Bulls and Cows
 * SecretNums = bull + (m - n) + cows
 */

// @lc code=start
func getHint(secret string, guess string) string {
	bull := 0
	// secret num
	hash := make([]int, 10)
	for i := 0; i < len(secret); i++ {
		indexS, _ := strconv.Atoi(string(secret[i]))
		indexG, _ := strconv.Atoi(string(guess[i]))
		if indexG == indexS {
			bull++
		} else {
			hash[indexS]++
			hash[indexG]--
		}
	}
	notenough := 0
	for i := 0; i < 10; i++ {
		if hash[i] > 0 {
			notenough = notenough + hash[i]
		}
	}
	cows := len(secret) - notenough - bull
	return fmt.Sprintf("%dA%dB", bull, cows)
}

// @lc code=end
