package linkedList_tree

import (
	"fmt"
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

func TestZigzagLevelOrder(t *testing.T) {
	req:=&TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(zigzagLevelOrder(req))

}


func TestTest(t *testing.T) {
	a:=[]int{1,2,3}
	fmt.Println(a)
}
