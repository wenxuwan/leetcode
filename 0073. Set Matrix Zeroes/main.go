/*
用一个数组来保存0的位置，后面再根据行列把其它置0
*/

package main

import "fmt"

func setZeroes(matrix [][]int)  {
	row := 0
	column := 0
	len1 := len(matrix[0])
	len2 := len(matrix)
	fmt.Println(len1,len2)
	zeroIndex := make([]int,0,len1*len2)
	for row < len2{
		for column < len1{
			if matrix[row][column] == 0{
				zeroIndex = append(zeroIndex, row)
				zeroIndex = append(zeroIndex, column)
			}
			column++
		}
		column = 0
		row++
	}
	for index := 0; index < len(zeroIndex);index = index + 2{
		for i := 0; i < len1;i++{
			matrix[zeroIndex[index]][i] = 0
		}
		for j := 0; j < len2;j++{
			matrix[j][zeroIndex[index + 1]] = 0
		}
	}
}
