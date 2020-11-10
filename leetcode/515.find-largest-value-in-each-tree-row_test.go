package leetcode

import (
	"reflect"
	"testing"
)

func Test_largestValues(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	Node9 := &TreeNode{Val: 9}
	Node3 := &TreeNode{Val: 3}
	Node5 := &TreeNode{Val: 5}
	Node3_1 := &TreeNode{Val: 3, Left: Node5, Right: Node3}
	Node2 := &TreeNode{Val: 2, Right: Node9}
	Node1 := &TreeNode{Val: 1, Left: Node3_1, Right: Node2}
	Node22 := &TreeNode{Val: 2}
	Node23 := &TreeNode{Val: 3}
	Node21 := &TreeNode{Val: 1, Left: Node22, Right: Node23}
	Node31 := &TreeNode{Val: 1}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: Node1,
			},
			want: []int{1, 3, 9},
		},
		{
			name: "2",
			args: args{
				root: Node21,
			},
			want: []int{1, 3},
		},
		{
			name: "3",
			args: args{
				root: Node31,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestValues(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
