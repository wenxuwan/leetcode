package main

import "fmt"

func isValidPos(board [][]byte, iStr byte, row, col int) bool {
	bigRow, bigCol := row/3, col/3
	for i := 0; i < 9; i++ {
		if board[row][i] == iStr && i != col || (board[i][col] == iStr && i != row) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+3*bigRow][j+3*bigCol] == iStr && (i+3*bigRow != row || j+3*bigCol != col) {
				return false
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if !isValidPos(board, board[i][j], i, j) {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	data := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(data))
}
