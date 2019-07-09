/*
该题目主要是找出置换的规律即可
1.行的两侧变量对调
2.按照对角线置换
*/
package main

import "fmt"

func rotate(matrix [][]int)  {
	row := len(matrix)
	column := len(matrix[0])
	//fmt.Println(matrix)
	for i := 0; i < row;i++{
		for j := 0;j < column / 2;j++{
			matrix[i][j],matrix[i][column - 1 - j] = matrix[i][column - 1 - j],matrix[i][j]
		}
	}
	//fmt.Println(matrix)

	for i := 0; i < row;i++{
		j := 0
		for j <= column - 1 -i{
			matrix[i][j],matrix[column - 1 -j][row - 1 - i] = matrix[column - 1 -j][row - 1 - i], matrix[i][j]
			j++
		}
	}
	fmt.Println(matrix)
}

func main(){
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(input)
}
