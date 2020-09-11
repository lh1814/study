package backtracking

import "fmt"

//import "fmt"

/*
@Author: by LH
@date:  2020/8/18
@function: 回溯
*/

//491.二进制手表 难度 简单
func readBinaryWatch(num int) []string {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	return readBinaryWatchdfs(n, num, 0, 0)
}
func readBinaryWatchdfs(n []int, nums, first, index int) (r []string) {
	if nums == first {
		x := 0
		for i := 0; i < nums; i++ {
			x ^= 1 << (n[i] - 1)
		}

		if x>>8 > 11 || x&255 > 59 {
			return
		}

		if (x & 255) == 0 {
			return []string{fmt.Sprintf("%d:00", x>>8)}
		}
		if x&255 < 10 {
			return []string{fmt.Sprintf("%d:0%d", x>>8, x&255)}
		}
		return []string{fmt.Sprintf("%d:%d", x>>8, x&255)}
	}

	for i := index; i < 12; i++ {
		n[i], n[first] = n[first], n[i]
		r = append(r, readBinaryWatchdfs(n, nums, first+1, i+1)...)
		n[first], n[i] = n[i], n[first]
	}

	return r
}
