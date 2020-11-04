package offer05

import "bytes"

/*
@Author: by LH
@date:  2020/10/28
@function:剑指 Offer 05. 替换空格

请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

	输入：s = "We are happy."
	输出："We%20are%20happy."
*/

func replaceSpace(s string) string {
	var buf bytes.Buffer

	for k := range s {
		if s[k] == byte(' ') {
			buf.WriteString("%20")
		} else {
			buf.WriteByte(s[k])
		}
	}
	return buf.String()
}
