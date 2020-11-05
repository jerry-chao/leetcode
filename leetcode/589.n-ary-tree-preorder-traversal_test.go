package leetcode

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	Node6 := &Node{Val: 6}
	Node5 := &Node{Val: 5}
	Node3 := &Node{Val: 3, Children: []*Node{Node5, Node6}}
	Node2 := &Node{Val: 2}
	Node4 := &Node{Val: 4}
	Node1 := &Node{Val: 1, Children: []*Node{Node3, Node2, Node4}}
	Node22 := &Node{Val: 2}
	Node26 := &Node{Val: 6}
	Node214 := &Node{Val: 14}
	Node212 := &Node{Val: 12}
	Node213 := &Node{Val: 13}
	Node210 := &Node{Val: 10}
	Node211 := &Node{Val: 11, Children: []*Node{Node214}}
	Node28 := &Node{Val: 8, Children: []*Node{Node212}}
	Node29 := &Node{Val: 9, Children: []*Node{Node213}}
	Node27 := &Node{Val: 7, Children: []*Node{Node211}}
	Node24 := &Node{Val: 4, Children: []*Node{Node28}}
	Node25 := &Node{Val: 5, Children: []*Node{Node29, Node210}}
	Node23 := &Node{Val: 3, Children: []*Node{Node26, Node27}}
	Node21 := &Node{Val: 1, Children: []*Node{Node22, Node23, Node24, Node25}}

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
			want: []int{1, 3, 5, 6, 2, 4},
		},
		{
			name: "1",
			args: args{
				root: Node21,
			},
			want: []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
