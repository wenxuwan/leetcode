package main

import "fmt"

func Min(A byte, B byte, C byte)byte{
	if A < B {
		if A < C{
			return A
		}else{
			return C
		}
	}else{
		if B < C{
			return B
		}else{
			return C
		}
	}
}

func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	if row == 0{
		return 0
	}
	column := len(matrix[0])
	var result [][]byte
	var maxSideLenth byte
	//fmt.Println(row, column)
	for i := 0;i < row; i++{
		result = append(result,make([]byte,column))
		result[i][0] = matrix[i][0]
		if matrix[i][0] > maxSideLenth{
			maxSideLenth = matrix[i][0]
		}
	}
	for j := 1; j < column;j++{
		result[0][j] = matrix[0][j]
		if matrix[0][j] > maxSideLenth{
			maxSideLenth = matrix[0][j]
		}
	}
	//fmt.Println(matrix)
	//fmt.Println(result)
	for i := 1; i < row; i++{
		for j := 1;j < column;j++{
			if matrix[i][j] == '1'{
				result[i][j] = Min(result[i][j - 1], result[i - 1][j - 1], result[i - 1][j]) + matrix[i][j] - '0'
			}else{
				result[i][j] = '0'
			}
			if result[i][j] > maxSideLenth{
				maxSideLenth = result[i][j]
			}
		}
	}
	fmt.Println(result)
	return int((maxSideLenth - '0') * (maxSideLenth - '0'))
}

func main(){
	//testData := [][]byte{{'1','0','1','0','0'}, {'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}
	testData := [][]byte{{'1','1','1','1'},{'1','1','1','1'},{'1','1','1','1'}}
	fmt.Println(maximalSquare(testData))
}
