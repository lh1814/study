package offer05

import "testing"

/*
@Author: by LH
@date:  2020/10/28
@function:
*/

func Test_replaceSpace(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "We are happy.",
			},
			want: "We%20are%20happy.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpace(tt.args.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
