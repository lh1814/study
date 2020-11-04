package offer13

import "testing"

/*
@Author: by LH
@date:  2020/11/2
@function:
*/

func Test_movingCount(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				m: 2,
				n: 3,
				k: 1,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				m: 3,
				n: 1,
				k: 0,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				x: 334,
			},
			want: 10,
		},
		{
			name: "2",
			args: args{
				x: 58,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.x); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
