package leetcode

import (
	"log"
	"math"
)

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	minMuationValue = math.MaxInt64
	visGenetic = map[string]bool{}
	dfsMinMutation(start, end, bank, 0)
	return minMuationValue
}

var geneticHash map[string][]string = map[string][]string{
	"A": {"C", "G", "T"},
	"C": {"A", "G", "T"},
	"G": {"A", "C", "T"},
	"T": {"A", "C", "G"},
}
var minMuationValue int
var visGenetic map[string]bool

func dfsMinMutation(start, end string, bank []string, mutation int) {
	log.Println(start, end)
	if start == end {
		if mutation < minMuationValue {
			minMuationValue = mutation
		}
		return
	}
	// divide
	for i := 0; i < len(start); i++ {
		for _, mutationStr := range geneticHash[string(start[i])] {
			// change
			newStart := string(start[:i]) + mutationStr + string(start[i+1:])
			visGenetic[newStart] = true
			if !visGenetic[newStart] && isInBank(newStart, bank) {
				dfsMinMutation(newStart, end, bank, mutation+1)
			}
		}
	}
}

func isInBank(str string, bank []string) bool {
	for i := 0; i < len(bank); i++ {
		if str == bank[i] {
			return true
		}
	}
	return false
}

// @lc code=end
