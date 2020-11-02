package leetcode

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	Node5 := &TreeNode{Val: 7}
	Node4 := &TreeNode{Val: 15}
	Node3 := &TreeNode{Val: 20, Left: Node4, Right: Node5}
	Node2 := &TreeNode{Val: 9}
	Node1 := &TreeNode{Val: 3, Left: Node2, Right: Node3}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: Node1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
