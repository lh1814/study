package dp_greedy

import (
	"fmt"
	"math"
	"sort"
)

/*
@Author: by LH
@date:  2020/8/4
@function:
*/

func pp(dp [][]int) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Printf("%d\t", dp[i][j])
		}
		fmt.Println()
	}
}

//221.最大正方形
func maximalSquare(matrix [][]byte) int {
	var max int
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = minInt(minInt(dp[i+1][j], dp[i][j+1]), dp[i][j]) + 1
				max = maxInt(dp[i+1][j+1], max)
			}
		}
	}

	return max * max
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func min(dp [][]int, i, j int) int {
	x := get(dp, i-1, j)
	y := get(dp, i, j-1)
	z := get(dp, i-1, j-1)
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}

func get(dp [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	return dp[i][j]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArea(matrix [][]byte, i, j int) int {
	//t := matrix[i][j]
	index := 0
	for index+i < len(matrix) && index+j < len(matrix[i]) {
		for x := 1; x <= index; x++ {
			if matrix[i][j+x] == 0 || matrix[i+x][j] == 0 || matrix[i+x][j+x] == 0 {
				return index
			}
		}
		index++
	}
	return index
}

//最大子序和
func maxSubArray(nums []int) int {
	var dp, max int
	max = math.MinInt64
	for i := 0; i < len(nums); i++ {
		dp = maxInt(dp+nums[i], nums[i])
		max = maxInt(dp, max)
	}

	return max
}

//三角形最小路径和
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([]int, len(triangle[len(triangle)-1])+1)

	for i := len(triangle) - 1; i > 0; i-- {
		for j := 0; j < len(triangle)-1; j++ {
			dp[j] = minInt(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

//俄罗斯套娃信封问题 hard
func maxEnvelopes(envelopes [][]int) int {
	var max int
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, len(envelopes))

	for i := 0; i < len(envelopes); i++ {
		var t int
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				t = maxInt(t, dp[j]+1)
			}
		}
		t = maxInt(t, 1)
		dp[i] = t
		max = maxInt(t, max)
	}

	return max
}

//300. 最长上升子序列 （上面问题的前置题）
func lengthOfLIS(nums []int) int {
	var max int
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		var t int
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				t = maxInt(t, dp[j]+1)
			}
		}
		dp[i] = maxInt(t, 1)
		max = maxInt(max, t)
	}
	fmt.Println(dp)
	return max
}

//有效的utf8
func validUtf8_2(data []int) bool {
	if len(data) == 0 {
		return false
	}
	//第一个0的位置 n+1
	zeroIndex := 1
	for zeroIndex <= 4 && intIndexIsBit(data[0], zeroIndex) {
		zeroIndex++
	}
	if intIndexIsBit(data[0], zeroIndex) || zeroIndex == 2 {
		return false
	}
	//判断前n为1
	for i := 2; i < zeroIndex; i++ {
		if len(data) < i || !intIndexIsBit(data[i-1], 1) || intIndexIsBit(data[i-1], 2) {
			return false
		}
	}

	if zeroIndex == 1 && len(data) > 1 {
		return validUtf8(data[1:])
	}

	if zeroIndex != 1 && len(data) > zeroIndex-1 {
		return validUtf8(data[zeroIndex-1:])
	}
	return true
}

//判断数字位是否为1 index为从左到右最大为8
func intIndexIsBit(b, index int) bool {
	return 0 != b&(1<<(8-index))
}

//utf8
func validUtf8(data []int) bool {
	cnt := 0
	for _, v := range data {
		if cnt == 0 {
			if v>>5 == 0b110 {
				cnt = 1
			} else if v>>4 == 0b1110 {
				cnt = 2
			} else if v>>3 == 0b11110 {
				cnt = 3
			} else if v>>7 != 0 {
				return false
			}
		} else {
			if v>>6 != 0b10 {
				return false
			}
			cnt--
		}
	}
	return cnt==0
}
//z -= (z*z - x) / (2*z)  牛顿法
