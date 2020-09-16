package main

//这个题目还有很多地方是可以优化的，比如每次都去循环检查小块区域的数据。其实是可以放到map里面来操作的。
func solveSudoku(board [][]byte) {
	backTrack(board, 0)
}

func backTrack(board [][]byte, curStep int) bool {
	if curStep == 81 {
		return true
	}
	row, col := curStep/9, curStep%9
	if board[row][col] != '.' {
		return backTrack(board, curStep+1)
	}
	for i := 1; i <= 9; i++ {
		if !isValidPos(board, i, curStep) {
			continue
		}
		board[row][col] = '0' + byte(i)
		if backTrack(board, curStep+1) {
			return true
		}
		board[row][col] = '.'
	}
	return false
}

func isValidPos(board [][]byte, i, step int) bool {
	iStr := '0' + byte(i)
	row, col := step/9, step%9
	bigRow, bigCol := row/3, col/3
	for i := 0; i < 9; i++ {
		if board[row][i] == iStr || board[i][col] == iStr {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+3*bigRow][j+3*bigCol] == iStr {
				return false
			}
		}
	}
	return true
}

func main() {
	data := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(data)
}
