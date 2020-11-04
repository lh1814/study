package offer09

import "container/list"

/*
@Author: by LH
@date:  2020/10/28
@function:剑指 Offer 09. 用两个栈实现队列

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]

*/
type CQueue struct {
	append, remove *list.List
}

func Constructor() CQueue {
	return CQueue{
		append: list.New(),
		remove: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.append.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.remove.Len() == 0 {
		for this.append.Len() > 0 {
			this.remove.PushBack(this.append.Remove(this.append.Back()))
		}
	}

	if this.remove.Len() != 0 {
		res := this.remove.Back()
		this.remove.Remove(res)
		return res.Value.(int)
	}
	return -1
}
