package leetcode

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
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
		// 		start: "AACCGGTT",
		// 		end:   "AACCGGTA",
		// 		bank: []string{
		// 			"AACCGGTA",
		// 		},
		// 	},
		// 	want: 1,
		// },
		// {
		// 	name: "2",
		// 	args: args{
		// 		start: "AACCGGTT",
		// 		end:   "AAACGGTA",
		// 		bank: []string{
		// 			"AACCGGTA", "AACCGCTA", "AAACGGTA",
		// 		},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "3",
		// 	args: args{
		// 		start: "AAAAACCC",
		// 		end:   "AACCCCCC",
		// 		bank: []string{
		// 			"AAAACCCC", "AAACCCCC", "AACCCCCC",
		// 		},
		// 	},
		// 	want: 3,
		// },
		{
			name: "4",
			args: args{
				start: "AACCGGTT",
				end:   "AAACGGTA",
				bank: []string{
					"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA",
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
