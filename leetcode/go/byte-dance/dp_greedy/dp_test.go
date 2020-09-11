package dp_greedy

import (
	"fmt"
	"testing"
)

/*
@Author: by LH
@date:  2020/8/4
@function:
*/

func TestMaximalSquare(t *testing.T) {
	req := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	if maximalSquare(req) != 4 {
		t.Fatal("should be 4")
	}
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func TestMaxEnvelopes(t *testing.T) {
	req := [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}

	maxEnvelopes(req)

}

func TestValidUtf8(t *testing.T){
	if !validUtf8([]int{1}){
		t.Fatal("should be true")
	}

	if validUtf8([]int{235,140,4}){
		t.Fatal("should be false")
	}

	if !validUtf8([]int{228,189,160,229,165,189,13,10}){
		t.Fatal("should be true")
	}
}

func TestAAA(t *testing.T){
	fmt.Println(0xa==11)
	fmt.Printf("%3b\n",2)
}
