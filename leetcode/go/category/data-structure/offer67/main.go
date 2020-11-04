package offer67

import "math"

/*
@Author: by LH
@date:  2020/10/28
@function:
*/

func strToInt(str string) int {
	if str == "" {
		return 0
	}

	start := false

	flag := 1

	var ans int

	for k := range str {
		if ans == 0 && flag == 1 && !start && str[k] == ' ' {
			continue
		}
		if !start && str[k] == '-' {
			flag = -1
			start = true
			continue
		} else if !start && str[k] == '+' {
			start = true
			continue
		} else {
			return 0
		}
		if str[k] >= '0' && str[k] <= '9' {
			ans *= 10
			ans += int(str[k] - '0')
		} else {
			return 0
		}
		if ans*flag <= math.MinInt32 || ans*flag >= math.MaxInt32 {
			return ans * flag
		}
	}

}
