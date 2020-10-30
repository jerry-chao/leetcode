package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	Node4 := &ListNode{Val: 4}
	Node3 := &ListNode{Val: 0, Next: Node4}
	Node2 := &ListNode{Val: 2, Next: Node3}
	Node1 := &ListNode{Val: 3, Next: Node2}
	Node4.Next = Node2
	Node21 := &ListNode{Val: 1}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				head: Node1,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				head: Node21,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
