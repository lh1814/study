package offer10

/*
@Author: by LH
@date:  2020/10/29
@function:10- I. 斐波那契数列
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。



示例 1：

输入：n = 2
输出：1
示例 2：

输入：n = 5
输出：5
*/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var (
		first  = 0
		second = 1
	)
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}

/*
	剑指 Offer 10- II. 青蛙跳台阶问题
	一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。

	答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

	示例 1：

	输入：n = 2
	输出：2
	示例 2：

	输入：n = 7
	输出：21
	示例 3：

	输入：n = 0
	输出：1
	提示：

	状态转移方程
	f(n) = f(n-1) + f(n-2)

	0 <= n <= 100
*/
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	var (
		first  = 1
		second = 1
	)

	for i := 2; i <= n; i++ {
		first, second = second, (first+second)%1000000007
	}
	return second
}
