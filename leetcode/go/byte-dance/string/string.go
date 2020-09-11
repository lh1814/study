package string

import (
	//"go-example/library/time"
	"strings"
)

/*
@Author: by LH
@date:  2020/7/9
@function:
*/

//type Node struct {
//	//实际数
//	Val int
//	//目标数
//	Aim int
//}
//
//func (n *Node) IsOk() bool {
//	return n.Val >= n.Aim
//}
//func (n *Node) Dec() {
//	n.Val--
//}
//func (n *Node) Inc() {
//	n.Val++
//}

//字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	mp := make(map[uint8]int)
	for k := range s1 {
		mp[s1[k]] += 1
	}
	var l, r int
	var c uint8
	for r < len(s2) {
		c = s2[r]
		mp[c]--
		r++
		//c不是s1中字母
		for l < r && mp[c] < 0 {
			mp[s2[l]]++
			l++
		}
		if r-l == len(s1) {
			return true
		}
	}
	return false
}

type stringMultiply struct {
	len int
	val []byte
}

//func (s stringMultiply) Add()

//字符串相乘
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	r := make([]byte, len(num2)+len(num1))

	l1, l2 := len(num1), len(num2)
	for i := l1 - 1; i >= 0; i-- {
		n1 := num1[i]
		for j := l2 - 1; j >= 0; j-- {
			n2 := num2[j]
			x := (n1-'0')*(n2-'0') + r[i+j+1]
			index := i + j + 1
			r[i+j+1] = x % 10
			r[index-1] += x / 10
		}
	}
	if r[0] == 0 {
		r = r[1:]
	}
	for k := range r {
		r[k] += '0'
	}
	return string(r)
}

func reverseWords(s string) string {
	//defer time.Trace("reverseWords")()
	var l, r int
	x := []byte(s)
	for l < len(s) {
		for l < len(s) && s[l] == ' ' {
			l++
		}
		r = l
		for r < len(s) && s[r] != ' ' {
			r++
		}

		if l == r {
			break
		}

		for i := 0; i < (r-l)/2; i++ {
			x[l+i], x[r-i-1] = x[r-i-1], x[l+i]
		}

		l = r + 1
	}

	for i := 0; i < len(s)/2; i++ {
		x[i], x[len(s)-1-i] = x[len(s)-1-i], x[i]
	}
	blank := 0

	for i := 0; i < len(s); i++ {
		if x[i] == ' ' && (i == 0 || (i > 0 && x[i-1] == ' ')) {
			blank++
			continue
		}
		x[i-blank] = x[i]
	}

	if len(x) > 0 && x[len(x)-1] == ' ' && x[0] != ' ' {
		return string(x[:len(x)-blank-1])
	}

	return string(x[:len(x)-blank])
}

//复制IP地址(DFS 回溯算法)
func restoreIpAddresses(s string) []string {
	r := make([]string, 0)
	dfs([]string{}, s, &r)
	return r
}

func dfs(path []string, s string, r *[]string) {
	if len(path) == 3 {
		if validIP(s) {
			*r = append(*r, strings.Join(append(path, s), "."))
		}
		return
	}
	for i := 1; i <= 3 && i < len(s); i++ {
		t := s[:i]
		if validIP(t) {
			dfs(append(path, t), s[i:], r)
		}
	}
}
func validIP(ip string) bool {
	if len(ip) > 3 || (len(ip) == 3 && ip > "255") || (len(ip) != 1 && ip[0] == '0') {
		return false
	}
	return true
}

//简化路径
func simplifyPath(path string) string {
	//filepath.Clean()
	var stack []string

	var l, r int
	for l < len(path) {
		if path[l] == '/' {
			l++
			continue
		}
		r = l + 1
		for r < len(path) && path[r] != '/' {
			r++
		}

		switch path[l:r] {
		case ".":
			l = r
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, path[l:r])
		}
		l = r
	}

	return "/" + strings.Join(stack, "/")
}
