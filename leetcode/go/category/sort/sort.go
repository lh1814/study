package sort

/*
@Author: by LH
@date:  2020/10/26
@function:
*/

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
//平均时间复杂度	最坏时间复杂度	最好时间复杂度	空间复杂度	稳定性
//	O(n²)
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

//6.快速排序
//func QuickSort(nums []int)[]int{
//
//}
