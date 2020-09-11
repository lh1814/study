package leetcode

import (
	"container/list"
	"fmt"
	"sort"
)

/*
@Author: by LH
@date:  2020/8/5
@function: 每日练习
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//打家劫舍I
func robi(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := []int{nums[0], maxInt(nums[0], nums[1])}
	for i := 2; i < len(nums); i++ {
		dp[i%2] = maxInt(nums[i]+dp[i%2], dp[(i-1)%2])
	}

	return maxInt(dp[0], dp[1])
}

//打家劫舍II
func robii(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	return maxInt(robi(nums[1:]), robi(nums[:len(nums)-1]))
}

//打家劫舍III
func robiii(root *TreeNode) int {
	return maxIntSlice(dfs(root))
}

//dfs
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	return []int{root.Val + l[1] + r[1], maxIntSlice(l) + maxIntSlice(r)}
}

func maxIntSlice(x []int) int {
	if len(x) < 2 {
		return 0
	}
	if x[0] > x[1] {
		return x[0]
	}
	return x[1]
}

func restoreString(s string, indices []int) string {
	b := make([]byte, len(s))
	for i, v := range indices {
		b[v] = s[i]
	}
	return string(b)
}

//不同的二叉搜索树
//笛卡尔乘积可得一阶状态转移方程
//F(n)=n∑(i=1) F(i-1)*F(n-i)
//卡塔兰数可得二阶状态转移方程
//C(n+1)=2(2n+1)/n+2 * C(n)
//C(0)=1
func numTrees(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}

//生成所有二叉搜索树
//func generateTrees(n int) []*TreeNode {
//
//	for i := 1; i <= n; i++ {
//		x := &TreeNode{
//			Val: i,
//		}
//
//		for
//
//	}
//
//}

func pancakeSort(A []int) []int {
	B := make([]int, len(A))
	for i := 0; i < len(B); i++ {
		B[i] = i + 1
	}
	fmt.Println(A)
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	return B
}

//696. 计数二进制子串
func countBinarySubstrings(s string) int {
	var ptr, pre, res int

	for ptr < len(s) {
		start := ptr
		for ptr < len(s) && s[ptr] == s[start] {
			ptr++
		}
		res += minInt(pre, ptr-start)
		pre = ptr - start
	}
	return res
}
func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//有效的括号 2020.08.14
func isValid(s string) bool {
	l := list.New()

	for _, v := range s {
		switch v {
		case '{', '[', '(':
			l.PushFront(v)
		case '}':
			if l.Len() <= 0 || l.Front().Value.(rune) != '{' {
				return false
			}
			l.Remove(l.Front())
		case ')':
			if l.Len() <= 0 || l.Front().Value.(rune) != '(' {
				return false
			}
			l.Remove(l.Front())
		case ']':
			if l.Len() <= 0 || l.Front().Value.(rune) != '[' {
				return false
			}
			l.Remove(l.Front())
		default:
			return false
		}
	}

	if l.Len() > 0 {
		return false
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func increase(m []float64) bool {
	if len(m) <= 1 {
		return true
	}

	for i := len(m) - 1; i > 0; i-- {
		if m[i] <= m[i-1] {
			return false
		}
	}

	return true
}

//880. 索引处的解码字符串
func decodeAtIndex(S string, K int) string {
	var cur int
	var last string
	for k, v := range S {
		if v > '1' && v <= '9' {
			num := int(v - '0')

			//超过k
			if num*cur >= K {
				if K%cur == 0 {
					return last
				}
				return decodeAtIndex(S[:k], K%cur)
			}
			cur *= num
			continue
		}
		last = string(v)
		cur++
		if cur == K {
			return last
		}
	}
	return ""
}

//647. 回文子串
func countSubstrings(s string) int {
	var ans int

	var l, r int
	for k := range s {
		ans++

		l = k
		r = k + 1

		//判断偶数个，偶数个往右
		for l >= 0 && r < len(s) && s[l] == s[r] {
			ans++
			l--
			r++
		}

		l = k - 1
		r = k + 1
		//判断奇数个
		for l >= 0 && r < len(s) && s[l] == s[r] {
			ans++
			l--
			r++
		}
	}
	return ans
}

//529.扫雷游戏
func updateBoard(board [][]byte, click []int) [][]byte {
	if len(click) < 2 {
		return board
	}
	if isM(board, click[0], click[1]) {
		board[click[0]][click[1]] = 'X'
	} else {
		help529(board, click[0], click[1])
	}
	return board
}

func help529(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[x]) {
		return
	}

	if board[x][y] != 'E' {
		return
	}

	count := countM(board, x, y)
	if count > 0 {
		board[x][y] = '0' + uint8(count)
	} else {
		board[x][y] = 'B'
		help529(board, x-1, y-1)
		help529(board, x-1, y)
		help529(board, x-1, y+1)

		help529(board, x, y+1)
		help529(board, x, y-1)

		help529(board, x+1, y+1)
		help529(board, x+1, y)
		help529(board, x+1, y-1)
	}

	return
}

func countM(board [][]byte, x, y int) (count int) {

	if isM(board, x-1, y-1) {
		count++
	}

	if isM(board, x-1, y) {
		count++
	}
	if isM(board, x-1, y+1) {
		count++
	}
	if isM(board, x, y-1) {
		count++
	}
	if isM(board, x, y+1) {
		count++
	}
	if isM(board, x+1, y-1) {
		count++
	}
	if isM(board, x+1, y) {
		count++
	}
	if isM(board, x+1, y+1) {
		count++
	}

	return
}

func isM(board [][]byte, x, y int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[x]) {
		return false
	}
	return board[x][y] == 'M'
}

//230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	a := 0
	t := help230(root, k, &a)
	if t != nil {
		return t.Val
	}
	return 0
}

func help230(root *TreeNode, k int, cur *int) *TreeNode {
	if root == nil {
		return nil
	}
	l := help230(root.Left, k, cur)
	if l != nil {
		return l
	}
	*cur += 1
	if *cur == k {
		return root
	}
	r := help230(root.Right, k, cur)
	if r != nil {
		return r
	}

	return nil
}

//@date: 2020-08-24
//@info: 459. 重复的子字符串
func repeatedSubstringPattern(s string) bool {
	r := 1
	for r <= len(s)/2 {
		if len(s)%r != 0 {
			r++
			continue
		}
		var ans = true
		for i := 0; i < len(s)/r; i++ {
			if s[r*i:r*(i+1)] != s[:r] {
				ans = false
				break
			}
		}
		if ans {
			return ans
		}
		r++
	}
	return false
}

//@date:	2020-08-25
//@info:	491. 递增子序列
func findSubsequences(nums []int) [][]int {

	return nil
}

