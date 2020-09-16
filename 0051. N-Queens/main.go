package main

import (
	"fmt"
	"math"
)

var tmpResult [][][]byte

func validSet(canSet [][]byte, row, column, n int) bool {
	for i := 0; i < n; i++ {
		if canSet[row][i] == 'Q' || canSet[i][column] == 'Q' {
			return false
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if math.Abs(float64(i-row)) == math.Abs(float64(j-column)) {
				if canSet[i][j] == 'Q' {
					return false
				}
			}
		}
	}
	return true
}

func newCanset(newCanSet [][]byte, n int, canSet [][]byte) {
	for i := 0; i < n; i++ {
		newCanSet[i] = make([]byte, n)
		copy(newCanSet[i], canSet[i])
	}
}
func helper(canSet [][]byte, n int, m int, row int) bool {
	if n == 0 {
		return true
	}
	for i := row; i < m; i++ {
		for j := 0; j < m; j++ {
			if canSet[i][j] == '.' && validSet(canSet, i, j, m) {
				canSet[i][j] = 'Q'
				if !helper(canSet, n-1, m, i) {
					canSet[i][j] = '.'
				} else {
					tmp := make([][]byte, m)
					newCanset(tmp, m, canSet)
					tmpResult = append(tmpResult, tmp)
					canSet[i][j] = '.'
					continue
				}
			}
		}
	}
	return false
}

func solveNQueens(n int) [][]string {
	canSet := make([][]byte, n)
	tmpResult = nil
	for i := 0; i < n; i++ {
		canSet[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			canSet[i][j] = '.'
		}
	}
	helper(canSet, n, n, 0)
	result := make([][]string, len(tmpResult))
	for index, value := range tmpResult {
		for i := 0; i < n; i++ {
			s := string(value[i])
			result[index] = append(result[index], s)
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	solveNQueens(5)
}
