package offer30

import "container/list"

/*
@Author: by LH
@date:  2020/10/28
@function:定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.min();   --> 返回 -3.
	minStack.pop();
	minStack.top();      --> 返回 0.
	minStack.min();   --> 返回 -2.

*/

type MinStack struct {
	data, min *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: list.New(),
		min:  list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.data.PushBack(x)
	if this.min.Len() == 0 || this.min.Back().Value.(int) >= x {
		this.min.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	x := this.data.Remove(this.data.Back()).(int)
	if this.min.Back().Value.(int) == x {
		this.min.Remove(this.min.Back())
	}
}

func (this *MinStack) Top() int {
	if this.data.Len() > 0 {
		return this.data.Back().Value.(int)
	}
	return -1
}

func (this *MinStack) Min() int {
	if this.min.Len() > 0 {
		return this.min.Back().Value.(int)
	}
	return -1
}
