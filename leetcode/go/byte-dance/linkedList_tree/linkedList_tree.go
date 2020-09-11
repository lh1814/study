package linkedList_tree

/*
@Author: by LH
@date:  2020/7/28
@function:
*/

type ListNode struct {
	Val  int
	Next *ListNode
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

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var fast, slow *ListNode
	slow = head
	fast = head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
