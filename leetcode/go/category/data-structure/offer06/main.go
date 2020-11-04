package offer06

/*
@Author: by LH
@date:  2020/10/28
@function:剑指 Offer 06. 从尾到头打印链表
	输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

	输入：head = [1,3,2]
	输出：[2,3,1]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	return ans
}
