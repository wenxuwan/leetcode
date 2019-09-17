package main

import "fmt"

var re bool
func dfs(board [][]byte,row int, column int, index int,word string,used [][]int){
	if re == true{
		return
	}
	if index == len(word){
		re = true
		return
	}
	if row < 0 ||  column < 0 || row >= len(board) || column >= len(board[0]){
		return
	}
	if used[row][column] == 1{
		//fmt.Println("used",row,column)
		return
	}
	if board[row][column] == word[index]{

		used[row][column] = 1

		index++
		dfs(board,row - 1, column, index,word,used)
		if re != true{
			dfs(board,row + 1, column, index,word,used)
		}
		if re != true{
			dfs(board,row, column - 1, index,word,used)
		}
		if re != true{
			dfs(board,row, column + 1, index,word,used)
		}
		used[row][column] = 0
	}
}

func exist(board [][]byte, word string)bool{
	row := len(board)
	column := len(board[0])
	var used [][]int
	re = false
	for i := 0;i < row; i++{
		for j := 0;j < column;j++{
			used = append(used,make([]int,column))
		}
	}
	for i := 0; i < row;i++{
		for j := 0;j < column;j++{
			//fmt.Println("start---------------")
			if board[i][j] == word[0]{
				dfs(board, i, j, 0, word, used)
			}
			if re == true{
				goto LABEL
			}
		}
	}
	LABEL:
	return re
}

func main(){
	word := string("aaaaaaaaaaaaaaaaaaaa")
	//board := [][]byte{{'B','B','C','E'},{'S','F','C','S'},{'B','D','E','E'}}
	board2 := [][]byte{{'a','a','a','a'},{'a','a','a','a'},{'a','a','a','a'},{'a','a','a','a'},{'a','a','a','b'}}
	fmt.Println(exist(board2,word))
}