package arr_sort

import (
	"fmt"
	"testing"
)

/*
@Author: by LH
@date:  2020/7/10
@function:
*/

func TestMaxAreaOfIsland(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
	//test(grid)
	//fmt.Println(grid)
}

func TestFindLengthOfLCIS(t *testing.T) {
	fmt.Println(findLengthOfLCIS([]int{1, 1, 1}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))
}

func TestMinimumTotal(t *testing.T) {
	p := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(p))
}

func TestLongestConsecutive(t *testing.T) {
	req := []int{0, 0, 1}
	if longestConsecutive(req) != 2 {
		t.Fatal("!=222222")
	}
}

func TestGetPermutation(t *testing.T) {
	if getPermutation(3, 3) != "213" {
		t.Fatal("err", getPermutation(3, 3))
	}

	if getPermutation(4, 9) != "2314" {
		t.Fatal("err")
	}
}

func TestTest(t *testing.T) {
	a := map[int]bool{}
	test(a)
	fmt.Println(a)
}

func test(a map[int]bool) {
	a[0] = true
}

func TestFindCircleNum(t *testing.T) {
	req := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	if findCircleNum(req) != 2 {
		t.Fatal(findCircleNum(req), "≠", 2)
	}

	req = [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 1},
	}

	if findCircleNum(req) != 1 {
		t.Fatal(findCircleNum(req), "≠", 1)
	}

	req = [][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}

	if findCircleNum(req) != 1 {
		t.Fatal(findCircleNum(req), "≠", 1)
	}

}
func TestMerge(t *testing.T) {

	a := [][]int{
		{1, 6},
		{2, 5},
		{0, 3},
		{8, 9},
		{4, 6},
	}
	fmt.Println(merge(a))

	a = [][]int{
		{1, 4},
		{0, 2},
		{3, 5},
	}
	fmt.Println(merge(a))

}
func TestTrap(t *testing.T) {
	var height []int
	var aim int
	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	aim = 6
	if trap(height) != aim {
		t.Fatal("should be ", aim)
	}

	height = []int{9}
	aim = 0
	if trap(height) != aim {
		t.Fatal("should be ", aim)
	}

	height = []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	aim = 23
	if trap(height) != aim {
		t.Fatal("should be ", aim)
	}
}

func TestMergeSlice(t *testing.T) {
	fmt.Println(restoreString("art", []int{1, 0, 2}))
}

func TestAA(t *testing.T) {
	fmt.Println(minFlips("10111"))
}
