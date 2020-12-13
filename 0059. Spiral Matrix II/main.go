package main

import "fmt"

var result [][]int

func helper(value int, flag int, n int) {
	rowNum := len(result) - flag
	columnNum := len(result[0]) - flag
	row := flag
	column := flag
	if (n+1)/2 == flag {
		return
	}
	for ; column < columnNum; column++ {
		result[row][column] = value
		value++
	}
	row++
	column--
	for ; row < rowNum; row++ {
		result[row][column] = value
		value++
	}
	row--
	column--
	for ; column >= flag; column-- {
		result[row][column] = value
		value++
	}

	row--
	column = flag
	for ; row > flag; row-- {
		result[row][column] = value
		value++
	}
	flag++
	helper(value, flag, n)
}
func generateMatrix(n int) [][]int {
	result = make([][]int, n, n)
	for index, _ := range result {
		result[index] = make([]int, n, n)
	}
	helper(1, 0, n)
	return result
}
func main() {
	fmt.Println(generateMatrix(3))
}
