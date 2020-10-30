package leetcode

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	Node4 := &ListNode{Val: 4}
	Node3 := &ListNode{Val: 3, Next: Node4}
	Node2 := &ListNode{Val: 2, Next: Node3}
	Node1 := &ListNode{Val: 1, Next: Node2}
	Node21 := &ListNode{Val: 1}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				head: Node1,
			},
			want: Node2,
		},
		{
			name: "2",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "3",
			args: args{
				head: Node21,
			},
			want: Node21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
