package leetcode

import "testing"

func Test_isMonotonic(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				A: []int{1, 2, 2, 3},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				A: []int{6, 5, 4, 4},
			},
			want: true,
		},

		{
			name: "false3",
			args: args{
				A: []int{1, 3, 2},
			},
			want: false,
		},

		{
			name: "true4",
			args: args{
				A: []int{1, 2, 4, 5},
			},
			want: true,
		},
		{
			name: "true5",
			args: args{
				A: []int{1, 1, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
