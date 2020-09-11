package category

import (
	"fmt"
	"time"
)

/*
@Author: by LH
@date:  2020/8/10
@function: 分类：蓄水池抽样
*/

// 382. 链表随机节点
/**
 * Definition for singly-linked list.
 */
//空间复杂度O(n)思路是建立数组[]*ListNode，取随机数和长度取余，时间复杂度O(1)
//空间复杂度O(1),时间复杂度O(n)，蓄水池算法，先取n个数，从n+1开始
type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	var res, index int
	index = 1
	cur := this.head
	for cur != nil {
		if time.Now().UnixNano()%int64(index) == 0 {
			res = cur.Val
		}
		cur = cur.Next
		index++
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

//398. 随机数索引
type Solution1 struct {
	value []int
}

func Constructor1(nums []int) Solution1 {
	return Solution1{value: nums}
}

func (this *Solution1) Pick(target int) int {
	var res, index int
	index = 1
	for i := 0; i < len(this.value); i++ {
		if this.value[i] == target {
			if time.Now().UnixNano()%int64(index) == 0 {
				res = i
			}
			index++
		}
	}
	return res
}

//470.用rand7()求rand10()
func rand7() int {
	return int(time.Now().UnixNano()%7) + 1
}

func rand10() int {
	var res int = 49

	for res > 40 {
		a := rand7()
		b := rand7()
		res = a + (b-1)*7
	}

	return res%10 + 1
}

//130.被围绕的区域
func solve(board [][]byte) {
	if len(board)<=0{
		return
	}
	for j := 0; j < len(board[0]); j++ {
		dfs(board, 0, j)
		dfs(board, len(board)-1, j)
	}
	for i := 0; i < len(board); i++ {
		//fmt.Println(i, 0)
		dfs(board, i, 0)
		dfs(board, i, len(board[0])-1)
	}
	fmt.Println()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%s\t",string(board[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]=='A'{
				board[i][j]='O'
			}else if board[i][j]=='O'{
				board[i][j]='X'
			}
		}
	}
}


func dfs(board [][]byte, i, j int) {
	fmt.Println(i,j)
	if i < 0 || i >= len(board) || j >= len(board[i]) || j < 0 || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	fmt.Println(";;;ll")

	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}

func bfs(board [][]byte, i, j int) {
	if i < 0 || i > len(board) || j < len(board[i]) || j < 0 || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
}


//478.在圆内随机生成点