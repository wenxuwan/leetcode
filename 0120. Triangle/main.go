package main

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	rowLen := len(triangle)
	result := math.MaxInt32
	if rowLen < 2{
		result = triangle[0][0]
	}
	for row := 1; row < rowLen;row++{
		l := len(triangle[row])
		triangle[row][0] += triangle[row - 1][0]
		triangle[row][l - 1] += triangle[row - 1][l - 2]
		for column := 1; column < l - 1;column++{
			if triangle[row - 1][column] > triangle[row - 1][column - 1]{
				triangle[row][column] += triangle[row - 1][column - 1]
			}else{
				triangle[row][column] += triangle[row - 1][column]
			}
		}
	}
	for i := 0; i < rowLen;i++{
		if result > triangle[rowLen - 1][i]{
			result = triangle[rowLen - 1][i]
		}
	}
	return result
}

func main(){
	test := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	fmt.Println(minimumTotal(test))
}
