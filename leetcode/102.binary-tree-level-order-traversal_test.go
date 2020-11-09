package leetcode

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
<<<<<<< HEAD
	type args struct {
		root *TreeNode
	}
	Node15 := &TreeNode{Val: 15}
	Node7 := &TreeNode{Val: 7}
	Node9 := &TreeNode{Val: 9}
	Node20 := &TreeNode{Val: 20, Left: Node15, Right: Node7}
	Node3 := &TreeNode{Val: 3, Left: Node9, Right: Node20}
=======
	Node7 := &TreeNode{Val: 7}
	Node15 := &TreeNode{Val: 15}
	Node20 := &TreeNode{Val: 20, Left: Node15, Right: Node7}
	Node9 := &TreeNode{Val: 9}
	Node3 := &TreeNode{Val: 3, Left: Node9, Right: Node20}
	type args struct {
		root *TreeNode
	}

>>>>>>> modify leetcode to project
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: Node3,
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
