package main

import "fmt"

func f(board [][]byte, flag_board [][]bool, word string, lenth int){

}

func exist(board [][]byte, word string)bool{
	fmt.Println(board)
	fmt.Println(word)
	fmt.Println(len(word))
	fmt.Println(board[0][0] == word[0])
	return true
}

func main(){
	word := string("ABCCED")
	board := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	exist(board,word)
}