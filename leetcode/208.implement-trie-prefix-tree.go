package leetcode

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start

// Trie describe Trie
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// TrieConstructor construct trie
func TrieConstructor() Trie {
	return Trie{}
}

// Insert Inserts a word into the trie.
func (trie *Trie) Insert(word string) {
	curr := trie
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

// Search Returns if the word is in the trie.
func (trie *Trie) Search(word string) bool {
	curr := trie
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

// StartsWith Returns if there is any word in the trie that starts with the given prefix.
func (trie *Trie) StartsWith(prefix string) bool {
	curr := trie
	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}
