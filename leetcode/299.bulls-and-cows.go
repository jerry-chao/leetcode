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
/* func getHint(secret string, guess string) string {
	buckets := make([]int, 10)
	bull := 0

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bull++
			continue
		}
		s, _ := strconv.Atoi(string(secret[i]))
		g, _ := strconv.Atoi(string(guess[i]))

		buckets[s]++
		buckets[g]--
	}

	sum := 0
	for _, bucket := range buckets {
		if bucket > 0 {
			sum = sum + bucket
		}
	}
	cows := len(secret) - sum - bull
	return fmt.Sprintf("%dA%dB", bull, cows)

} */

func getHint(secret string, guess string) string {
	bull := 0
	hash := make([]int, 10)
	for i := 0; i < 10; i++ {
		hash[i] = 0
	}
	for i := 0; i < len(secret); i++ {
		if guess[i] == secret[i] {
			bull++
		}
		secretIndex, _ := strconv.Atoi(string(secret[i]))
		hash[secretIndex]++
		guessIndex, _ := strconv.Atoi(string(guess[i]))
		hash[guessIndex]--
	}
	sum := 0
	for _, num := range hash {
		if num > 0 {
			sum = sum + num
		}
	}
	cows := len(secret) - bull - sum
	return fmt.Sprintf("%dA%dB", bull, cows)
}

// @lc code=end
