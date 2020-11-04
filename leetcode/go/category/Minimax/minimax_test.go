package Minimax

import (
	"testing"
)

/*
@Author: by LH
@date:  2020/9/11
@function:
*/

func Test_getMoneyAmount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{5},
			want: 6,
		},
		{
			name: "test2",
			args: args{8},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMoneyAmount(tt.args.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
