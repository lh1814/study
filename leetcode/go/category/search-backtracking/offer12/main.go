package offer12

/*
@Author: by LH
@date:  2020/11/2
@function:剑指 Offer 12. 矩阵中的路径

请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

[["a","b","c","e"],
["s","f","c","s"],
["a","d","e","e"]]

但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。

示例 1：

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：

输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false
*/

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if help(board, word, i, j) == true {
				return true
			}
		}
	}
	return false
}

func help(board [][]byte, word string, x, y int) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[x]) || len(word) == 0 {
		return false
	}
	tmp := board[x][y]
	if tmp != word[0] || tmp == ' ' {
		return false
	}
	if len(word) == 1 {
		return true
	}
	board[x][y] = ' '
	defer func() {
		board[x][y] = tmp
	}()

	return help(board, word[1:], x+1, y) || help(board, word[1:], x, y+1) || help(board, word[1:], x-1, y) || help(board, word[1:], x, y-1)
}
