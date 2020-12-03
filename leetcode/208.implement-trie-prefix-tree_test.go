package leetcode

import "testing"

func TestTrie_Insert(t *testing.T) {
	type fields struct {
		children [26]*Trie
		isEnd    bool
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := &Trie{
				children: tt.fields.children,
				isEnd:    tt.fields.isEnd,
			}
			trie.Insert(tt.args.word)
		})
	}
}
