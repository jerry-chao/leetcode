package leetcode

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	visMutations = make(map[int]bool, len(bank)+1)
	minMutationValue = len(bank) + 1
	dfsMutation(start, end, bank, 0)
	if minMutationValue <= len(bank) {
		return minMutationValue
	}
	return -1
}

var mutations = map[uint8][]string{
	'A': {"C", "G", "T"},
	'C': {"A", "G", "T"},
	'G': {"A", "C", "T"},
	'T': {"A", "C", "G"},
}

var visMutations map[int]bool
var minMutationValue int

func dfsMutation(start, end string, bank []string, mutaion int) {
	// terminator
	if mutaion >= minMutationValue {
		return
	}
	if start == end {
		minMutationValue = mutaion
	}
	// divide problem
	for i := 0; i < len(start); i++ {
		for _, mutationValue := range mutations[start[i]] {
			newStart := string(start[:i]) + mutationValue + string(start[i+1:])
			if idx := indexOf(newStart, bank); idx != -1 && !visMutations[idx] {
				visMutations[idx] = true
				dfsMutation(newStart, end, bank, mutaion+1)
				visMutations[idx] = false
			}
		}
	}
	// relate sub problem
	// revert
}

func indexOf(str string, bank []string) int {
	for i := 0; i < len(bank); i++ {
		if str == bank[i] {
			return i
		}
	}
	return -1
}

// @lc code=end
