package leetcode

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "1",
		// 	args: args{
		// 		nums: []int{2, 3, -2, 4},
		// 	},
		// 	want: 6,
		// },
		// {
		// 	name: "2",
		// 	args: args{
		// 		nums: []int{-2, 0, -1},
		// 	},
		// 	want: 0,
		// },
		{
			name: "3",
			args: args{
				nums: []int{-4, -3, -2},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
