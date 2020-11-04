package _1365

/*
@Author: by LH
@date:  2020/10/26
@function:
*/
//@info: 1365. 有多少小于当前数字的数字
func smallerNumbersThanCurrent(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				ans[j]++
			}
		}
	}
	return ans
}
