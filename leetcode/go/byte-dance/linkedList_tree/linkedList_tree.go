package linkedList_tree

import "container/list"

/*
@Author: by LH
@date:  2020/7/28
@function:
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r, tmp *ListNode
	var over int

	for l1 != nil || l2 != nil || over > 0 {
		if r == nil {
			tmp = &ListNode{}
			r = tmp
		} else {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}

		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		tmp.Val += over
		over = tmp.Val / 10
		tmp.Val %= 10

	}

	return r
}

//排序链表
func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	t := head
	var length int
	for t != nil {
		length++
		t = t.Next
	}

	for i := 1; i < length; i <<= 1 {
		cur := dummy.Next
		tail := dummy

		for cur != nil {
			left := cur
			right := cutList(left, i)
			cur = cutList(right, i)
			tail.Next = mergeTwoLists(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummy.Next
}

func cutList(head *ListNode, n int) *ListNode {
	var tmp *ListNode
	for n > 0 && head != nil {
		n--
		tmp = head
		head = head.Next
	}
	if tmp != nil {
		tmp.Next = nil
	}
	return head
}

//合并有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	////递归
	//if l1 == nil {
	//	return l2
	//}
	//if l2 == nil {
	//	return l1
	//}
	//if l1.Val > l2.Val {
	//	l2.Next = mergeTwoLists(l1, l2.Next)
	//	return l2
	//}
	//l1.Next = mergeTwoLists(l1.Next, l2)
	//return l1

	head := &ListNode{}
	tmp := head
	//迭代
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
		} else {
			tmp.Next = l1
			l1 = l1.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else if l2 == nil {
		tmp.Next = l1
	}

	return head.Next
}

//环形链表II
func detectCycle(head *ListNode) *ListNode {
	var fast, slow *ListNode
	slow = head
	fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

//相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	//我的呆萌写法
	//var lenA, lenB int
	//testA := headA
	//for testA != nil && testA.Next != nil {
	//	testA = testA.Next
	//	lenA++
	//}
	//testB := headB
	//for testB != nil && testB.Next != nil {
	//	testB = testB.Next
	//	lenB++
	//}
	//
	//if lenA > lenB {
	//	for i := 0; i < lenA-lenB; i++ {
	//		headA = headA.Next
	//	}
	//} else {
	//	for i := 0; i < lenB-lenA; i++ {
	//		headB = headB.Next
	//	}
	//}
	//for headA != nil && headB != nil {
	//	if headA == headB {
	//		return headA
	//	}
	//	headA = headA.Next
	//	headB = headB.Next
	//}

	//题解区大佬写法
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

//合并K个排序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) > 2 {
		m := len(lists) / 2
		return mergeTwoLists(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
	} else if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	} else if len(lists) == 1 {
		return lists[0]
	}
	return nil
}

//二叉树的最近公共祖先
//思路 前序遍历
//递归 回溯
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//基线条件
	if root == p || root == q || root == nil {
		return root
	}

	//递归条件
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	if r != nil {
		return r
	}

	return nil
}

//二叉树的锯齿形层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	l := list.New()
	isOrder := true
	dummy := &TreeNode{}
	l.PushFront(dummy)
	l.PushFront(root)
	var tmp []int
	for l.Len() > 0 {
		cur := l.Front().Value.(*TreeNode)
		l.Remove(l.Front())
		if cur != dummy {
			if isOrder {
				tmp = append(tmp, cur.Val)
			} else {
				tmp = append([]int{cur.Val}, tmp...)
			}

			if cur.Left != nil {
				l.PushBack(cur.Left)
			}
			if cur.Right != nil {
				l.PushBack(cur.Right)
			}
		} else {
			res = append(res, tmp)
			tmp = nil
			isOrder = !isOrder
			if l.Len() > 0 {
				l.PushBack(dummy)
			}
		}
	}
	return res
}
