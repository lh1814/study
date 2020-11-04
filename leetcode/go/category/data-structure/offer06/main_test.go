package offer06

import (
	"reflect"
	"testing"
)

/*
@Author: by LH
@date:  2020/10/28
@function:
*/

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
