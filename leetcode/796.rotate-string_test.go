package leetcode

import "testing"

func Test_rotateString(t *testing.T) {
	type args struct {
		A string
		B string
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
				A: "abcde",
				B: "cdeab",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				A: "abcde",
				B: "abced",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateString(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("rotateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
