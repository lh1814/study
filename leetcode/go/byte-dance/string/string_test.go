package string

import (
	"fmt"
	"testing"
)

/*
@Author: by LH
@date:  2020/7/9
@function:
*/

func TestCheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	if !checkInclusion(s1, s2) {
		t.Fatal("should be true")
	}
	s1 = "ab"
	s2 = "eidboaoo"
	if checkInclusion(s1, s2) {
		t.Fatal("should be false")
	}
}
func TestMultiply(t *testing.T) {
	//var a byte
	fmt.Print(multiply("102", "99"))
}

func TestReverseWords(t *testing.T) {

	fmt.Println(reverseWords("   "))
}

func TestRestoreIpAddresses(t *testing.T) {
	//fmt.Println("1111"[4:])
	fmt.Println(restoreIpAddresses("0000"))
}
func TestSimplifyPath(t *testing.T) {
	//if simplifyPath("/home/")!="/home"{
	//	t.Fatal("not /home")
	//}
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
