package leetcode

import "testing"

func Test_getKthMagicNumber(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{k: 5},
			want: 9,
		},
		{
			name: "2",
			args: args{k: 7},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthMagicNumber(tt.args.k); got != tt.want {
				t.Errorf("getKthMagicNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
