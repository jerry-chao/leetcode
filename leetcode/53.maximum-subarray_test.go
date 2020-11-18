package leetcode

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "5",
			args: args{
				nums: []int{-2147483647},
			},
			want: -2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
