package linkedList_tree

import (
	"reflect"
	"testing"
)

/*
@Author: by LH
@date:  2020/7/28
@function:
*/

func TestSortList(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	res := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	if !reflect.DeepEqual(res, sortList(head)) {
		t.Fatal("TestSortList should be equal")
	}
}
