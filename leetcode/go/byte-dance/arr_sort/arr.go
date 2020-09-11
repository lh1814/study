package arr_sort

import (
	"bytes"
	"sort"
)

/*
@Author: by LH
@date:  2020/7/10
@function:
*/

//岛屿最大面积
func maxAreaOfIsland(grid [][]int) int {
	var r int
	h := len(grid)
	for i := 0; i < h; i++ {
		w := len(grid[i])
		for j := 0; j < w; j++ {
			res := 0
			help(grid, h, w, i, j, &res)
			if res > r {
				r = res
			}
		}
	}
	return r
}
func help(x [][]int, h, w, i, j int, value *int) {

	if i < 0 || i >= h || j < 0 || j >= w || x[i][j] == 0 {
		return
	}
	*value++
	x[i][j] = 0
	help(x, h, w, i+1, j, value)
	help(x, h, w, i, j+1, value)
	help(x, h, w, i-1, j, value)
	help(x, h, w, i, j-1, value)
	return
}

//最长连续递增序列 (滑动窗口)
func findLengthOfLCIS(nums []int) int {
	var l, r, max int
	for l < len(nums) {
		r++
		for r < len(nums) && nums[r] > nums[r-1] {
			r++
		}
		if r-l > max {
			max = r - l
		}
		l = r
	}

	return max

	//var l, r, max int
	////r = 1
	//for r < len(nums) {
	//	if r < len(nums) && nums[r] > nums[r-1] {
	//		r++
	//		continue
	//	}
	//	if r-l > max {
	//		max = r - l
	//	}
	//	l = r
	//}
	//if r-l > max {
	//	max = r - l
	//	l = r + 1
	//}
	//return max
}

//三角形最小路径和
func minimumTotal(triangle [][]int) int {
	//n := len(triangle)
	//dp := make([][]int, n+1)
	//for k := range dp {
	//	dp[k] = make([]int, n+1)
	//}
	//for i := n - 1; i >= 0; i-- {
	//	for j := 0; j <= i; j++ {
	//		dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
	//	}
	//}
	//return dp[0][0]

	n := len(triangle)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//func help2(triangle [][]int, i, j int) int {
//
//	if i >= len(triangle) {
//		return 0
//	}
//	return 0
//
//}

//数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	if k >= len(nums) {
		return -1
	}
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//排序算法的十种实现

//1.冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return
}

//2.选择排序
func SelectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		t := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[t] {
				t = j
			}
		}
		nums[i], nums[t] = nums[t], nums[i]
	}
}

//3.插入排序
func InsertionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		t := i
		for j := i - 1; j >= 0; j-- {
			if nums[t] < nums[j] {
				nums[t], nums[j] = nums[j], nums[t]
				t = j
			}
		}
	}
}

//4.希尔排序
func ShellSort(nums []int) {

}

//5.归并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	m := len(nums) / 2
	return mergeSlice(MergeSort(nums[:m]), MergeSort(nums[m:]))
}

//合并切片
func mergeSlice(left, right []int) []int {
	r := make([]int, len(left)+len(right))
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			r[i+j] = left[i]
			i++
		} else {
			r[i+j] = right[j]
			j++
		}
	}
	for i < len(left) {
		r[i+j] = left[i]
		i++
	}
	for j < len(right) {
		r[i+j] = right[j]
		j++
	}
	return r
}

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}

	for _, num := range nums {
		numSet[num] = true
	}

	var long int
	for k := range numSet {
		if !numSet[k-1] {
			t := k + 1
			for numSet[t] {
				t++
			}
			if t-k > long {
				long = t - k
			}
		}
	}
	return long
}

//第k个排列
func getPermutation(n int, k int) string {
	var buf bytes.Buffer
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = i + 1
	}

	var t = k - 1 //余
	for i := 0; i < n; i++ {
		d := factorial(n - i - 1)
		index := t / d
		t %= d
		buf.WriteByte('0' + uint8(tmp[index]))
		tmp = deleteItem(tmp, index)
	}
	return buf.String()
}

func factorial(n int) int {
	var r = 1
	for n > 1 {
		r *= n
		n--
	}
	return r
}

func deleteItem(t []int, index int) []int {
	return append(t[:index], t[index+1:]...)
}

//朋友圈个数
func findCircleNum(M [][]int) int {
	var r int
	//看题解是用数组相比map好一点，此处可换
	tmp := make(map[int]bool)
	for i := 0; i < len(M); i++ {
		if tmp[i] {
			continue
		}
		r++
		tmp[i] = true
		helpCircleNum(M, i, tmp)
	}
	return r
}

//将上下左右值变0
func helpCircleNum(m [][]int, i int, tmp map[int]bool) {
	for j := 0; j < len(m[i]); j++ {
		if !tmp[j] && m[i][j] == 1 {
			tmp[j] = true
			helpCircleNum(m, j, tmp)
		}
	}
}

//合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	index := 0
	i := 1
	for i < len(intervals) {
		//需要合并
		if intervals[index][1] >= intervals[i][0] {
			//intervals[index][0] = intervals[i-1][0]
			if intervals[index][1] < intervals[i][1] {
				intervals[index][1] = intervals[i][1]
			}
			i++
			continue
		}
		//不合并
		index++
		intervals[index] = intervals[i]
		i++
	}
	return intervals[:index+1]
}

//接雨水
func trap(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		var tmp, max int
		//左指针移动
		if height[left] < height[right] {
			max = left
			for left < len(height) && left < right && height[max] >= height[left] {
				tmp += height[left]
				left++
			}
			res += (left-max)*height[max] - tmp
		} else {
			max = right
			for right >= 0 && left < right && height[max] >= height[right] {
				tmp += height[right]
				right--
			}
			res += (max-right)*height[max] - tmp
		}
	}
	return res
}
func restoreString(s string, indices []int) string {
	r := make([]byte, len(s))
	for i := 0; i < len(indices); i++ {
		r[indices[i]] = s[i]
	}
	return string(r)
}

//1529. 灯泡开关 IV
func minFlips(target string) int {
	var num int
	var l, r int
	for l < len(target) {
		if target[l] == '0' {
			l++
			continue
		}
		r = l
		for r < len(target) && target[r] == '1' {
			r++
		}
		l = r
		num++
	}

	if target[len(target)-1] == '1' {
		return num*2 - 1
	}
	return num * 2

}

//1531. 压缩字符串 II
func getLengthOfOptimalCompression(s string, k int) int {
	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {

	return 0
}
