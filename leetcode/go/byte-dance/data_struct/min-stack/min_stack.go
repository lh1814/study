package min_stack

/*
@Author: by LH
@date:  2020/7/6
@function: 最小栈
*/

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if this.isEmpty() {
		this.stack = append(this.stack, 0)
		this.min = x
		return
	}
	interval := x - this.min
	if interval < 0 {
		this.min = x
	}
	this.stack = append(this.stack, interval)
}

func (this *MinStack) Pop() {
	if len(this.stack) <= 0 {
		return
	}
	if this.stack[len(this.stack)-1] < 0 {
		this.min = this.min - this.stack[len(this.stack)-1]
	}

	this.stack = this.stack[:len(this.stack)-1]

}

func (this *MinStack) Top() int {
	if len(this.stack) <= 0 {
		return -1
	}

	if this.stack[len(this.stack)-1] <= 0 {
		return this.min
	}

	return this.stack[len(this.stack)-1] + this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}

func (this *MinStack) isEmpty() bool {
	return len(this.stack) <= 0
}
