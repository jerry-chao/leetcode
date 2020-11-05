package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	Node5 := &ListNode{Val: 5}
	Node4 := &ListNode{Val: 4, Next: Node5}
	Node3 := &ListNode{Val: 3, Next: Node4}
	Node2 := &ListNode{Val: 2, Next: Node3}
	Node1 := &ListNode{Val: 1, Next: Node2}
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
			want: Node5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
