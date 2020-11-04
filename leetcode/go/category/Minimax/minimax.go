package Minimax

import "math"

/*
@Author: by LH
@date:  2020/9/11
@function:
*/

/**
 *@info: 375. 猜数字大小 II
 *@param: n
 */
func getMoneyAmount(n int) int {
	return getMoneyAmountForce(1, n)
}

//暴力解法
func getMoneyAmountForce(begin, end int) int {
	if begin >= end {
		return 0
	}
	var ans = math.MaxInt64
	for i := begin; i <= end; i++ {
		t := i + max(getMoneyAmountForce(begin, i-1), getMoneyAmountForce(i+1, end))
		ans = min(t, ans)
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
