package tencent

import (
	"fmt"
	"reflect"
	"testing"
)

/*
@Author: by LH
@date:  2020/8/14
@function:
*/

func TestMerge(t *testing.T) {
	//nums1:=[]int{1,2,3,0,0,0}
	//m := 3
	//nums2:=[]int{2,5,6}
	//n:=3
	//merge(nums1,m,nums2,n)
	//if !reflect.DeepEqual(nums1,[]int{1,2,2,3,5,6}){
	//	t.Fatal("nums1 is ",nums1)
	//}

	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(nums1, []int{1}) {
		t.Fatal("nums1 is ", nums1)
	}

}

func TestMajorityElement(t *testing.T) {

	//fmt.Println(majorityElement([]int{6,6,6,7,7}))
	//fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))

	//fmt.Println(reverseWords("Let's take LeetCode contest"))
	var a =&[]int{22,22,22}
	//*a = append(*a, 22)
	//*a = append(*a, 22)
	//*a = append(*a, 22)
	fmt.Println(a,len(*a),cap(*a),&(*a)[0])
	aa(a)
	fmt.Println(a,len(*a),cap(*a),&(*a)[0])
}

func aa(a *[]int){
	*a = append(*a, 33)
}

func TestGenerateMatrix(t *testing.T) {
	//0,0 0,1 0,2 ... 0,n 1,n 2,n ... n,n
	//PrintArr(generateMatrix(6))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
	//PrintArr()
}

func PrintArr(x [][]int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Printf("%d \t", x[i][j])
		}
		fmt.Println()
	}
}

func TestPermute(t *testing.T){
	nums:=[]int{1,2,3}
	PrintArr(permute(nums))
	fmt.Println(nums)
}