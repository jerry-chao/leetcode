package leetcode

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
<<<<<<< HEAD
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
=======
 *
 * https://leetcode.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (42.42%)
 * Likes:    470
 * Dislikes: 63
 * Total Accepted:    34.4K
 * Total Submissions: 80.8K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * A gene string can be represented by an 8-character long string, with choices
 * from "A", "C", "G", "T".
 *
 * Suppose we need to investigate about a mutation (mutation from "start" to
 * "end"), where ONE mutation is defined as ONE single character changed in the
 * gene string.
 *
 * For example, "AACCGGTT" -> "AACCGGTA" is 1 mutation.
 *
 * Also, there is a given gene "bank", which records all the valid gene
 * mutations. A gene must be in the bank to make it a valid gene string.
 *
 * Now, given 3 things - start, end, bank, your task is to determine what is
 * the minimum number of mutations needed to mutate from "start" to "end". If
 * there is no such a mutation, return -1.
 *
 * Note:
 *
 *
 * Starting point is assumed to be valid, so it might not be included in the
 * bank.
 * If multiple mutations are needed, all mutations during in the sequence must
 * be valid.
 * You may assume start and end string is not the same.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 *
 * return: 1
 *
 *
 *
 *
 * Example 2:
 *
 *
 * start: "AACCGGTT"
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * return: 2
 *
 *
 *
 *
 * Example 3:
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * return: 3
 *
 *
 *
 *
 */

// @lc code=start
// BFS
var mutationMap = map[uint8][3]string{
	'A': [...]string{"T", "G", "C"},
	'C': [...]string{"T", "G", "A"},
	'T': [...]string{"A", "G", "C"},
	'G': [...]string{"T", "A", "C"},
}

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func dfsMinMutation(curr string, end string, count int, bank []string) {
	// terminator
	if count >= minCount {
		return
	}
	if curr == end {
		if count < minCount {
			minCount = count
		}
		return
	}

	for i := 0; i < len(curr); i++ {
		for _, c := range mutationMap[curr[i]] {
			gene := curr[:i] + c + curr[i+1:]
			if idx := idxOf(gene, bank); idx != -1 && !isUsed[idx] {
				isUsed[idx] = true
				dfsMinMutation(gene, end, count+1, bank)
				isUsed[idx] = false
			}
		}
	}
}

var isUsed []bool
var minCount int

func minMutation(start string, end string, bank []string) int {
	// terminator
	if idxOf(end, bank) == -1 {
		return -1
	}
	minCount = len(bank) + 1
	isUsed = make([]bool, len(bank))
	dfsMinMutation(start, end, 0, bank)
	if minCount <= len(bank) {
		return minCount
>>>>>>> modify leetcode to project
	}
	return -1
}

// @lc code=end
