/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
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
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				matrix: [][]int{},
				target: 0,
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				matrix: [][]int{{}},
				target: 0,
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				matrix: [][]int{
					{1},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				matrix: [][]int{
					{1, 1},
				},
				target: 2,
			},
			want: false,
		},
		{
			name: "7",
			args: args{
				matrix: [][]int{
					{1, 3},
				},
				target: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf(" target: %v searchMatrix() = %v, want %v", tt.args.target, got, tt.want)
			}
		})
	}
}
