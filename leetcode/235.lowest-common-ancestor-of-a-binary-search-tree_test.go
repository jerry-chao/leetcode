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
	Node9 := &TreeNode{Val: 9}
	Node7 := &TreeNode{Val: 7}
	Node3 := &TreeNode{Val: 3}
	Node5 := &TreeNode{Val: 5}
	Node0 := &TreeNode{Val: 0}
	Node4 := &TreeNode{Val: 4, Left: Node3, Right: Node5}
	Node2 := &TreeNode{Val: 2, Left: Node0, Right: Node4}
	Node8 := &TreeNode{Val: 8, Left: Node7, Right: Node9}
	Node6 := &TreeNode{Val: 6, Left: Node2, Right: Node8}
	Node21 := &TreeNode{Val: 1}
	Node22 := &TreeNode{Val: 2, Left: Node21}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: Node6,
				p:    Node2,
				q:    Node8,
			},
			want: Node6,
		},
		{
			name: "2",
			args: args{
				root: Node6,
				p:    Node2,
				q:    Node4,
			},
			want: Node2,
		},
		{
			name: "3",
			args: args{
				root: Node22,
				p:    Node22,
				q:    Node21,
			},
			want: Node22,
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
