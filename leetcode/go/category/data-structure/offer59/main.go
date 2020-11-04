package offer59

import "container/list"

/*
@Author: by LH
@date:  2020/10/28
@function:剑指 Offer 59 - I. 滑动窗口的最大值
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]

*/

/*
	动态规划解法、
	1.按k划分区
	2.建立两数组分别记录 区块从左到坐标点最大值、区块从右到坐标点最大值
	3.比较left[i] right[i+k-1]最大值
*/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var ans []int
	queue := list.New()
	clean := func(i int) {
		if queue.Len() > 0 && queue.Front().Value.(int) <= i-k {
			queue.Remove(queue.Front())
		}
		for queue.Len() > 0 && nums[queue.Back().Value.(int)] <= nums[i] {
			queue.Remove(queue.Back())
		}
	}
	for i := 0; i < k; i++ {
		clean(i)
		queue.PushBack(i)
	}
	ans = append(ans, nums[queue.Front().Value.(int)])
	for i := k; i < len(nums); i++ {
		clean(i)
		queue.PushBack(i)
		ans = append(ans, nums[queue.Front().Value.(int)])
	}
	return ans
}
