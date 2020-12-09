package main

import "fmt"

var result []int

func helper(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	rowNum := len(matrix)
	columnNum := len(matrix[0])
	row := 0
	column := 0
	if rowNum == 1 || columnNum == 1 {
		for row = 0; row < rowNum; row++ {
			for column = 0; column < columnNum; column++ {
				result = append(result, matrix[row][column])
			}
		}
		return
	}
	for column = 0; column < columnNum; column++ {
		result = append(result, matrix[row][column])
	}
	column--
	for row = 1; row < rowNum; row++ {
		result = append(result, matrix[row][column])
	}
	row--
	column--
	for ; column >= 0; column-- {
		result = append(result, matrix[row][column])
	}
	column = 0
	row--
	for ; row > 0; row-- {
		result = append(result, matrix[row][column])
	}
	newMatrix := matrix[1 : rowNum-1]
	for index, value := range newMatrix {
		newMatrix[index] = value[1 : columnNum-1]
	}
	helper(newMatrix)
}
func spiralOrder(matrix [][]int) []int {
	result = make([]int, 0, len(matrix)*len(matrix[0]))
	helper(matrix)
	return result
}

func main() {
	array := [][]int{{1, 2}, {3, 4}, {5, 6}, {5, 6}, {5, 6}, {5, 6}, {5, 6}, {5, 6}, {5, 6}, {5, 6}}
	spiralOrder(array)
	fmt.Println(result)
}
