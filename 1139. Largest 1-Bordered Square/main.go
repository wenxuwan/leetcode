package main

import "fmt"

func largest1BorderedSquare(grid [][]int) int {
	if grid == nil {
		return 0
	}
	result := 0
	rowLen := len(grid)
	columnLen := len(grid[0])
	tmpGrid := make([][][]int, len(grid)+2, len(grid)+2)
	for index, _ := range tmpGrid {
		tmpGrid[index] = make([][]int, len(grid[0])+2, len(grid[0])+2)
		for index2, _ := range tmpGrid[index] {
			tmpGrid[index][index2] = make([]int, 2, 2)
		}
	}
	for row := 0; row < rowLen; row++ {
		for column := 0; column < columnLen; column++ {
			if grid[row][column] == 1 {
				tmpGrid[row+1][column+1][0] = tmpGrid[row+1][column][0] + 1 //向左
				tmpGrid[row+1][column+1][1] = tmpGrid[row][column+1][1] + 1 //向上
			}
			l := 1
			if tmpGrid[row+1][column+1][0] > tmpGrid[row+1][column+1][1] {
				l = tmpGrid[row+1][column+1][1]
			} else {
				l = tmpGrid[row+1][column+1][0]
			}
			for l > 0 {
				if tmpGrid[row+1][column+2-l][1] >= l && tmpGrid[row+2-l][column+1][0] >= l {
					if l*l > result {
						result = l * l
					}
				}
				l--
			}
		}
	}
	return result
}

func main() {
	data := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	fmt.Println(largest1BorderedSquare(data))
}
