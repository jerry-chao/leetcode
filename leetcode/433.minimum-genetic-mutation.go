package leetcode

import "log"

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	minMuationValue = len(bank) + 1
	visGenetic = make([]bool, len(bank))
	dfsMinMutation(start, end, bank, 0)
	if minMuationValue <= len(bank) {
		return minMuationValue
	}
	return -1
}

var geneticHash map[string][]string = map[string][]string{
	"A": {"C", "G", "T"},
	"C": {"A", "G", "T"},
	"G": {"A", "C", "T"},
	"T": {"A", "C", "G"},
}
var minMuationValue int
var visGenetic []bool

func dfsMinMutation(start, end string, bank []string, mutation int) {
	if mutation >= minMuationValue {
		return
	}
	if start == end {
		if mutation < minMuationValue {
			minMuationValue = mutation
		}
		return
	}
	// divide
	log.Println("start", start, "end:", end, mutation, minMuationValue)
	for i := 0; i < len(start); i++ {
		for _, mutationStr := range geneticHash[string(start[i])] {
			// change
			newStart := start[:i] + mutationStr + start[i+1:]
			if idx := idxInBank(newStart, bank); idx != -1 && !visGenetic[idx] {
				visGenetic[idx] = true
				dfsMinMutation(newStart, end, bank, mutation+1)
				visGenetic[idx] = false
			}
		}
	}
}

func idxInBank(str string, bank []string) int {
	for i := 0; i < len(bank); i++ {
		if str == bank[i] {
			return i
		}
	}
	return -1
}

// @lc code=end
