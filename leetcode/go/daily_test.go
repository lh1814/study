package leetcode

import (
	"fmt"
	"testing"
	"time"
)

/*
@Author: by LH
@date:  2020/8/5
@function:
*/

func TestRob(t *testing.T) {
	//if robi([]int{1,2,3,1})!=4{
	//	t.Fatal("should be 4")
	//}

	if robii([]int{2, 3, 2}) != 3 {
		t.Fatal("should be 3")
	}
	if robii([]int{1, 2, 3, 1}) != 4 {
		t.Fatal("should be 4:", robii([]int{1, 2, 3, 1}))
	}
}
func TestRobIII(t *testing.T) {
	req := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Println(dfs(req))

	if robiii(req) != 7 {
		t.Fatal("should be 7:", robiii(req))
	}
}

func TestIsValid(t *testing.T) {
	if !isValid("()") {
		t.Fatal("() should be true")
	}

	if !isValid("(){}[]") {
		t.Fatal("(){}[] should be true")
	}

	if isValid("(]") {
		t.Fatal("(] should be false")
	}

	if isValid("([)]") {
		t.Fatal("([)] should be false")
	}

	if !isValid("{[]}") {
		t.Fatal("{[]} should be true")
	}
}

func TestArr(t *testing.T) {
	//up :=[]float64{3.1,3.2,3.3,3.4}
	//stable:=[]float64{3.1,3.0,3.3,3.4}
	//decline := []float64{3.5,3.4,3.3,3.2}
	//
	//
	//fmt.Println(status(up))
	//fmt.Println(status(stable))
	//fmt.Println(status(decline))

	//fmt.Println(decodeAtIndex("leet2code3",10))
	//fmt.Println(decodeAtIndex("ha22",5))
	//fmt.Println(decodeAtIndex("a2345678999999999999999999999999999999999999999999999999999999999999 ",1))

	fmt.Println(countSubstrings("aaa"))

}
func PrintArr(x [][]byte) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Printf("%c \t", x[i][j])
		}
		fmt.Println()
	}
}
func TestUpdateBoard(t *testing.T) {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	click := []int{3, 0}
	PrintArr(updateBoard(board, click))
	fmt.Println()

	board = [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'M', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}
	click = []int{1, 2}
	PrintArr(updateBoard(board, click))
}

func TestKthSmallest(t *testing.T) {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(kthSmallest(root, 5))

}

func TestRepeatedSubstringPattern(t *testing.T) {
	fmt.Println(repeatedSubstringPattern("abab"))
}

func p(c chan int) {
	fmt.Println(c)
	time.Sleep(time.Second * 1)
	c <- 33
	fmt.Println("ok")

}

func TestAA(t *testing.T) {

	//c := make(chan int)
	//fmt.Println(",", c)
	//go p(c)
	//time.Sleep(time.Second * 2)
	//fmt.Println(<-c)

}

func TestFindSubsequences(t *testing.T) {
	a :=[]int{1,2}
	fmt.Println(a[2:],2/2)
}

//func TestLetterCombinations