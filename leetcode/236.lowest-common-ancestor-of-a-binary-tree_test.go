package leetcode

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	Node8 := &TreeNode{Val: 8}
	Node0 := &TreeNode{Val: 0}
	Node1 := &TreeNode{Val: 1, Left: Node0, Right: Node8}
	Node4 := &TreeNode{Val: 4}
	Node7 := &TreeNode{Val: 7}
	Node2 := &TreeNode{Val: 2, Left: Node7, Right: Node4}
	Node6 := &TreeNode{Val: 6}
	Node5 := &TreeNode{Val: 5, Left: Node6, Right: Node2}
	Node3 := &TreeNode{Val: 3, Left: Node5, Right: Node1}
	Node22 := &TreeNode{Val: 2}
	Node21 := &TreeNode{Val: 1, Left: Node22}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: Node3,
				p:    Node5,
				q:    Node1,
			},
			want: Node3,
		},
		{
			name: "2",
			args: args{
				root: Node3,
				p:    Node5,
				q:    Node4,
			},
			want: Node5,
		},
		{
			name: "3",
			args: args{
				root: Node21,
				p:    Node21,
				q:    Node22,
			},
			want: Node21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
