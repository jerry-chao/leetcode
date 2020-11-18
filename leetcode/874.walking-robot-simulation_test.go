package leetcode

import "testing"

func Test_robotSim(t *testing.T) {
	type args struct {
		commands  []int
		obstacles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				commands:  []int{4, -1, 3},
				obstacles: [][]int{},
			},
			want: 25,
		},
		{
			name: "2",
			args: args{
				commands:  []int{4, -1, 4, -2, 4},
				obstacles: [][]int{{2, 4}},
			},
			want: 65,
		},
		{
			name: "3",
			args: args{
				commands:  []int{4, -1, 4, -2, 4},
				obstacles: [][]int{{2, 4}},
			},
			want: 65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim(tt.args.commands, tt.args.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}
