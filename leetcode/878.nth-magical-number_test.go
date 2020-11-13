package leetcode

import "testing"

func Test_nthMagicalNumber(t *testing.T) {
	type args struct {
		N int
		A int
		B int
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
				N: 1,
				A: 2,
				B: 3,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				N: 4,
				A: 2,
				B: 3,
			},
			want: 6,
		},
		{
			name: "3",
			args: args{
				N: 5,
				A: 2,
				B: 4,
			},
			want: 10,
		},
		{
			name: "4",
			args: args{
				N: 3,
				A: 6,
				B: 4,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthMagicalNumber(tt.args.N, tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("nthMagicalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
