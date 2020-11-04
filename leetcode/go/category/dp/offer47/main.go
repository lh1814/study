package offer47

/*
@Author: by LH
@date:  2020/10/29
@function:剑指 Offer 47.礼物的最大价值
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

示例 1:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
*/

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 {
				if j == 0 {
					continue
				} else {
					grid[i][j] += grid[i][j-1]
				}
			} else {
				if j == 0 {
					grid[i][j] += grid[i-1][j]
				} else {
					grid[i][j] += Max(grid[i][j-1], grid[i-1][j])
				}
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
