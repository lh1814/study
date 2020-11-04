package offer58

import "testing"

/*
@Author: by LH
@date:  2020/10/28
@function:
*/

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		n int
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "abcdefg",
				n: 2,
			},
			want: "cdefgab",
		},
		{
			name: "2",
			args: args{
				s: "lrloseumgh",
				n: 6,
			},
			want: "umghlrlose",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
