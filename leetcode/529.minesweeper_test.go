package leetcode

import (
	"reflect"
	"testing"
)

func Test_updateBoard(t *testing.T) {
	type args struct {
		board [][]byte
		click []int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				board: [][]byte{
					[]byte("EEEEE"),
					[]byte("EEMEE"),
					[]byte("EEEEE"),
					[]byte("EEEEE"),
				},
				click: []int{3, 0},
			},
			want: [][]byte{
				[]byte("B1E1B"),
				[]byte("B1M1B"),
				[]byte("B111B"),
				[]byte("BBBBB"),
			},
		},
		{
			name: "2",
			args: args{
				board: [][]byte{
					[]byte("B1E1B"),
					[]byte("B1M1B"),
					[]byte("B111B"),
					[]byte("BBBBB"),
				},
				click: []int{1, 2},
			},
			want: [][]byte{
				[]byte("B1E1B"),
				[]byte("B1X1B"),
				[]byte("B111B"),
				[]byte("BBBBB"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBoard(tt.args.board, tt.args.click); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
