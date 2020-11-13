package leetcode

import "math"

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (30.41%)
 * Likes:    3971
 * Dislikes: 1332
 * Total Accepted:    488.2K
 * Total Submissions: 1.6M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * the length of shortest transformation sequence from beginWord to endWord,
 * such that:
 *
 *
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list.
 *
 *
 * Note:
 *
 *
 * Return 0 if there is no such transformation sequence.
 * All words have the same length.
 * All words contain only lowercase alphabetic characters.
 * You may assume no duplicates in the word list.
 * You may assume beginWord and endWord are non-empty and are not the same.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * Output: 5
 *
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog",
 * return its length 5.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * Output: 0
 *
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 *
 *
 *
 *
 *
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordID = map[string]int{}
	graph = [][]int{}
	// add word to graph
	for _, word := range wordList {
		addEdge(word)
	}
	beginID := addEdge(beginWord)
	endID, has := wordID[endWord]
	if !has {
		return 0
	}

	// init all word to maxInt
	const inf int = math.MaxInt64
	dist := make([]int, len(wordID))
	for i := range dist {
		dist[i] = inf
	}

	dist[beginID] = 0
	queue := []int{beginID}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endID {
			return dist[endID]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

var wordID map[string]int
var graph [][]int

func addWord(word string) int {
	if had, ok := wordID[word]; ok {
		return had
	}
	id := len(wordID)
	wordID[word] = id
	graph = append(graph, []int{})
	return id
}

func addEdge(word string) int {
	id := addWord(word)
	s := []byte(word)
	for i, b := range s {
		s[i] = '*'
		otherID := addWord(string(s))
		graph[id] = append(graph[id], otherID)
		graph[otherID] = append(graph[otherID], id)
		s[i] = b
	}
	return id
}

// @lc code=end
