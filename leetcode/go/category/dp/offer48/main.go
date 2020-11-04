package offer48

/*
@Author: by LH
@date:  2020/10/29
@function:剑指 Offer 48. 最长不含重复字符的子字符串
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。



示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring(s string) int {
	exist := make(map[byte]int)
	var (
		left, right, max int
	)

	for right < len(s) {
		index, ok := exist[s[right]]
		if ok && index >= left {
			//存在移动左坐标,判定最大值
			max = Max(max, right-left)
			left = index + 1
		}
		exist[s[right]] = right
		right++
	}
	max = Max(max, right-left)

	return max
}
