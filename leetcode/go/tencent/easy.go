package tencent

import (
	"bytes"
)

/*
@Author: by LH
@date:  2020/8/14
@function:
*/

//88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[index] = nums1[m]
			m--
			index--
		} else {
			nums1[index] = nums2[n]
			n--
			index--
		}
	}

	for i := m; i >= 0; i-- {
		nums1[i] = nums1[i]
	}
	for i := n; i >= 0; i-- {
		nums1[i] = nums2[i]
	}
}

//169. 多数元素
func majorityElement(nums []int) int {
	//排序法
	//sort.Ints(nums)
	//return nums[len(nums)/2]

	//摩尔投票法
	var candidate, count int

	for _, v := range nums {

		if count <= 0 {
			candidate = v
			count = 1
		} else {
			if candidate != v {
				count--
			} else {
				count++
			}
		}
	}

	return candidate
}

//557. 反转字符串中的单词 III
func reverseWords(s string) string {
	var buf bytes.Buffer

	var l, r int
	for l < len(s) {
		if s[l] == ' ' {
			l++
			continue
		}
		r = l
		for r < len(s) && s[r] != ' ' {
			r++
		}
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		for i := r - 1; i >= l; i-- {
			buf.WriteByte(s[i])
		}

		l = r
	}

	return buf.String()
}

//292. Nim 游戏
func canWinNim(n int) bool {
	return n%4 != 0
}

//231. 2的幂
func isPowerOfTwo(n int) bool {
	//1
	//var x int = 1
	//for x < n {
	//	x *= 2
	//}
	//return x == n

	if n <= 0 {
		return false
	}

	//return n&-n == 0

	return n&(n-1) == 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

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

//59. 螺旋矩阵 II
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	t := 1
	x := 0
	for t <= n*n {

		for i := x; i < n-x && t <= n*n; i++ {
			ans[x][i] = t
			t++
		}

		for i := x + 1; i < n-x && t <= n*n; i++ {
			ans[i][n-x-1] = t
			t++
		}

		for i := n - x - 2; i >= x && t <= n*n; i-- {
			ans[n-x-1][i] = t
			t++
		}

		for i := n - x - 2; i > x && t <= n*n; i-- {
			ans[i][x] = t
			t++
		}

		x += 1
	}
	return ans
}

//78. 子集
func subsets(nums []int) [][]int {
	ans := [][]int{{}}

	for _, v := range nums {
		x := make([][]int, 0)
		//fmt.Println("ans",ans)
		for _, v1 := range ans {
			//fmt.Println("tt",v1)
			x = append(x, append([]int{v}, v1...))
			//fmt.Println(v,v1,append(v1, v))
		}
		ans = append(ans, x...)
	}
	return ans
}

//46.全排列
func permute(nums []int) [][]int {
	return permutedfs(nums, 0)
}

func permutedfs(nums []int, first int) (r [][]int) {
	if first == len(nums) {
		t:=make([]int,len(nums))
		copy(t,nums)
		r = append(r, t)
		return
	}

	for i := first; i < len(nums); i++ {
		nums[first], nums[i] = nums[i], nums[first]
		r = append(r, permutedfs(nums, first+1)...)
		nums[i], nums[first] = nums[first], nums[i]
	}
	return
}

